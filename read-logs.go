package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("~/.npm/_logs")
	file, err := os.Open("2020-04-20T09_38_57_076Z-debug.log")
	if err != nil {
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
