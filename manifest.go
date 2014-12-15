package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
)
import "fmt"

type File_data struct {
	Name string
	Hash string
}

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

			file_obj := File_data{current[size:len(current)], hash_str}
			file_json, _ := json.Marshal(file_obj)
			fmt.Println(string(file_json))

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

	fmt.Println(iter(path, len(path)))

}
