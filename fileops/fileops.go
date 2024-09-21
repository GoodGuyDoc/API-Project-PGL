package fileops

import (
	"fmt"
	"os"
)

// This function is used to for basic file writing. if the file does not exist it will create one.
// TODO use an interface type so that we can convert any type to string in the function

func WriteToFile(filename string, data string) error {
	//Types must first be converted to a string (if not already) then to a byte for writing file
	err := os.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		return err
	} else {
		fmt.Printf("File written to %s\n", filename)
	}
	return nil
}

func ReadDataFile(filename string) (data string, err error) {
	dataRead, err := os.ReadFile(filename)
	data = string(dataRead) //Must convert into proper string data
	if err != nil {
		return "", err
	}
	//From string data the dev can convert it to whatever format they need.
	return data, nil
}
