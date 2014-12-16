package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type File_data struct {
	Name     string
	Hash     string
	Buildnum uint8
}

func obj2byte(obj *File_data) []byte {
	obj_byte := make([]byte, 0)
	obj_byte = append(obj_byte, uint8(len(obj.Name)))
	obj_byte = append(obj_byte, []byte(obj.Name)...)
	obj_byte = append(obj_byte, uint8(20)) //This is static, otherwise the base conversion would fuck this up
	obj_byte = append(obj_byte, []byte(obj.Hash)...)
	obj_byte = append(obj_byte, obj.Buildnum)

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

			file_obj := File_data{current[size:len(current)], string(hash_raw), 0}

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

func parse_iter(obj []byte) {
	name_size := obj[0]
	name := string(obj[1 : name_size+1])
	hash_size := obj[name_size+1]
	hash := hex.EncodeToString(obj[name_size+2 : hash_size+2+name_size])
	buildnum := obj[hash_size+2+name_size]
	file_obj := File_data{name, hash, buildnum}

	fmt.Println(file_obj)

	if obj[hash_size+3+name_size] != 127 {
		parse_iter(obj[hash_size+3+name_size : len(obj)])
	}
}

func parse(obj []byte) uint8 {
	if string([]byte{0xCA, 0xFE, 0xBA, 0xBE}) != string(obj[0:4]) {
		fmt.Println("Invalid magic number")
		return 0
	} else {
		if obj[len(obj)-1] != 127 {
			fmt.Println("ESC missing")
			return 0
		} else {

			if obj[4] == 127 {
				fmt.Println("Warning, zero data in manifest")
				return 1
			}

			parse_iter(obj[4:len(obj)])
			return 1
		}
	}
}

func main() {
	fmt.Println("qTox updater manifest tool")

	if len(os.Args) < 2 {
		fmt.Println("No flags specified")
		return
	}

	opt_gen := flag.Bool("generate", false, "Generate manifest")
	opt_val := flag.Bool("validate", false, "Validate manifest")
	opt_manifest := flag.String("manifest", "manifest", "Path to a manifest")
	opt_directory := flag.String("directory", "payload", "Path to a directory")

	flag.Parse()

	if *opt_gen && *opt_val {
		fmt.Println("Error, can't generate a manifest and validate a manifest at the same time")
		return
	} else if *opt_gen {

		path := *opt_directory + "/"
		obj_list = []byte{0xCA, 0xFE, 0xBA, 0xBE}

		if iter(path, len(path)) == -1 {
			fmt.Println("No folder found")
			return
		}

		obj_list = append(obj_list, 127)

		ioutil.WriteFile(*opt_manifest, obj_list, 0644)

		fmt.Println("Saved manifest")

	} else {
		manifest, _ := ioutil.ReadFile(*opt_manifest)
		if parse(manifest) != 1 {
			fmt.Println("validation failed")

		}
	}
}
