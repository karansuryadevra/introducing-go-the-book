package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // should close the file when all the surrounding functions are complete

	stat, err := file.Stat() // get file info
	if err != nil {
		panic(err)
	}

	fmt.Println(stat)

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs) // Read the file
	if err != nil {
		panic(err)
	}
	str := string(bs)
	fmt.Println(str)

	// A quicker and better way to read a file :-
	bsNew, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	strNew := string(bsNew)
	fmt.Println(strNew)

	// create a file

	newFile, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile.WriteString("This is a new file!")

	// Open a directory
	dir, err := os.Open("../../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // -1 means return the entire number of entries
	if err != nil {
		panic(err)
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// walk a file path

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
