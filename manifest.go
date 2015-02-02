package main

import (
	"bytes"
	"crypto/sha1"
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

var File_cache = make(map[string]File_data)

func obj2byte(obj *File_data) []byte {
	obj_byte := make([]byte, 0)
	//You got that James Dean day dream look in your eye
	obj_byte = append(obj_byte, uint8(len(obj.Name)))
	//And I got that red lip classic thing that you like
	obj_byte = append(obj_byte, uint8(20)) //[1]
	//And when we go crashing down, we come back every time.
	obj_byte = append(obj_byte, []byte(obj.Name)...)
	//Cause we never go out of style
	obj_byte = append(obj_byte, []byte(obj.Hash)...)
	//We never go out of style
	obj_byte = append(obj_byte, obj.Buildnum)
	obj_byte = append(obj_byte, 0x80)

	return obj_byte
}

//[1] This is static, otherwise the base conversion would fuck this up

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

			filename := current[size:len(current)]
			buildnum := uint8(0)

			if File_cache[filename].Name != "" {
				if File_cache[filename].Hash != string(hash_raw) {
					buildnum = File_cache[filename].Buildnum + uint8(1)
				}
			} else {
				fmt.Println(File_cache[filename].Name)
			}

			file_obj := File_data{current[size:len(current)], string(hash_raw), buildnum}

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
	hash_size := obj[1]

	str_len := int(name_size) + int(hash_size) + 2

	if str_len > len(obj) {
		fmt.Println("Corrupt data passed")
		return
	}

	if obj[str_len+1] != 128 {
		fmt.Println("Corrupt data passed")
		return
	}

	name := string(obj[2 : name_size+2])
	hash := string(obj[name_size+2 : hash_size+2+name_size])
	buildnum := obj[hash_size+2+name_size]

	file_obj := File_data{name, hash, buildnum}

	File_cache[name] = file_obj

	if obj[hash_size+4+name_size] != 127 {
		parse_iter(obj[hash_size+4+name_size : len(obj)])
	}
}

func parse(obj []byte) uint8 {

	if len(obj) < 5 {
		fmt.Println("Corrupt data passed")
		return 0
	}
	if string([]byte{0xCA, 0xFE, 0xBA, 0xBE}) != string(obj[0:4]) {
		fmt.Println("Invalid magic number")
		return 0
	} else {
		if obj[len(obj)-21] != 127 {
			fmt.Println("ESC missing")
			return 0
		} else {

			if obj[4] == 127 {
				fmt.Println("Warning, zero data in manifest")
				return 1
			}

			hash := sha1.New()
			hash.Write([]byte(obj[0 : len(obj)-20]))
			hash_raw := hash.Sum(nil)
			hash_old_raw := obj[len(obj)-20:]

			if bytes.Equal(hash_old_raw, hash_raw) != true {
				fmt.Println("Hash invalid")
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

		if _, err := os.Stat(*opt_manifest); err == nil {
			manifest, _ := ioutil.ReadFile(*opt_manifest)
			fmt.Println("Manifest found, updating")
			if parse(manifest) != 1 {
				fmt.Println("validation failed")
			}
		} else {
			fmt.Println("Manifest file not found, making a new one")
		}

		if iter(path, len(path)) == -1 {
			fmt.Println("No folder found")
			return
		}

		obj_list = append(obj_list, 0x7F)

		hash := sha1.New()
		hash.Write([]byte(obj_list))
		hash_raw := hash.Sum(nil)
		obj_list = append(obj_list, hash_raw...)

		ioutil.WriteFile(*opt_manifest, obj_list, 0644)

		fmt.Println("Saved manifest")

	} else {
		manifest, _ := ioutil.ReadFile(*opt_manifest)
		if parse(manifest) != 1 {
			fmt.Println("validation failed")
		} else {
			fmt.Println("Loaded manifest to global hashtable")
			for Name, Num := range File_cache {
				fmt.Println("Name:", Name, "Version:", Num.Buildnum)
			}
		}
	}
}
