/*
Defer is used to ensure that a function call is performed later in a program’s executio. This will be executed at the end of the enclosing function (main).
Defer is often used where e.g. ensure and finally would be used in other languages.
*/

/*
panic - A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation
A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	newFile := CreateFile("/go.txt")
	defer CloseFile(newFile) // This will ensure closing the file, when terminating main function
	WriteFile(newFile)
}

func CreateFile(path string) *os.File {
	newFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return newFile
}

func WriteFile(file *os.File) {
	fmt.Println("Inside write function")
	_, err := file.WriteString("Hello program")
	if err != nil {
		panic(err)
	}
	fmt.Println("write function is done")
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("file closed successfully")
}
