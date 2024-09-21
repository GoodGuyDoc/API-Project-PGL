package main

import (
	"API-Project-PGL/fileops"
	"fmt"
)

func main() {
	err := fileops.WriteToFile("testfile.txt", "Hello, World!")
	if err != nil {
		fmt.Println(err)
	}

	fileData, err := fileops.ReadDataFile("testfile.txt")
	if err != nil {
		fmt.Println("There was an issue reading the file")
	} else {
		fmt.Printf("The file reads %v", fileData)
	}

}
