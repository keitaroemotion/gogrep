package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)	

func getArgs() []string {
	args := os.Args[1:]
	if(len(args) == 0) {
		panic("at least 1 argument is required!")
	}
	return args
}

//
// TODO: multithread execution
//
func checkFiles(dir_path string, pattern string) {
	files, err := ioutil.ReadDir(dir_path); check(err)

	for _, file := range files {
		if file.IsDir() {
			checkFiles(file.Name(), pattern)
		} else {
			checkFile (file.Name(), pattern) 
		}
	}
}

func checkFile(file_path string, pattern string) {
	file_content, err := os.ReadFile(file_path); check(err)

	// This pattern matching part is not that easy.
	// Given the pattern includes regex this program
	// should parse/convert those regex and match accordingly
	// (maybe one-liner) not sure.

	result := fmt.Sprintf("[%s] %t", 
		file_path,
		strings.Contains(string(file_content), pattern))
	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args    := getArgs()
	pattern := args[0]
	checkFiles(".", pattern)
}
