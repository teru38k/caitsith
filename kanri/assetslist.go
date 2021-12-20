package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var ls []string
	files, err := ioutil.ReadDir("./image")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		ls = append(ls, file.Name())
	}
	fmt.Println(ls)
}
