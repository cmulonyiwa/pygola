package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var jsonData interface{}

func getFile(name string) (*os.File, error) {
	//pass a file that contains a json data
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	file, err := getFile("simple.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewDecoder(file).Decode(&jsonData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonData)

}
