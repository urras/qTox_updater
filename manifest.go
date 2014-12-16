package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os"
)
import "fmt"

type File_data struct {
	Name     string
	Hash     string
	Buildnum string
}

func obj2byte(obj *File_data) []byte {
	obj_byte := make([]byte, 0)
	obj_byte = append(obj_byte, uint8(len(obj.Name)))
	obj_byte = append(obj_byte, []byte(obj.Name)...)
	obj_byte = append(obj_byte, uint8(len(obj.Hash)))
	obj_byte = append(obj_byte, []byte(obj.Hash)...)
	obj_byte = append(obj_byte, uint8(len(obj.Buildnum)))
	obj_byte = append(obj_byte, []byte(obj.Buildnum)...)

	return obj_byte
}

var obj_list []byte

func fs_type(path string) int {
	f, err := os.Open(path)
	if err != nil {
		return -1
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return -1
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return 0
	case mode.IsRegular():
		return 1
	}

	return -1
}

func iter(path string, size int) int {
	if fs_type(path) != 0 {
		return -1
	}

	files, _ := ioutil.ReadDir(path)

	for _, file := range files {
		current := path + file.Name()

		if fs_type(current) == 1 {
			f, _ := ioutil.ReadFile(current)
			hash := sha1.New()

			hash.Write([]byte(f))
			hash_raw := hash.Sum(nil)
			hash_str := hex.EncodeToString(hash_raw)

			file_obj := File_data{current[size:len(current)], hash_str, "0"}

			obj_list = append(obj_list, obj2byte(&file_obj)...)

		} else if fs_type(current) == 0 {
			if iter(current+"/", size) != 1 {
				return -1
			}

		} else {
			return -1

		}
	}

	return 1
}

func main() {

	path := "payload/"
	obj_list = []byte{0xC, 0xA, 0xF, 0xE, 0xB, 0xA, 0xB, 0xE}

	if iter(path, len(path)) == -1 {
		fmt.Println("No folder found")
		return
	}
	obj_list = append(obj_list, 0x3)

	fmt.Println(obj_list)

}
