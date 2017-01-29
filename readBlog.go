package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("\nHello world:Start\n")

	files, _ := ioutil.ReadDir("html")
	for _, f := range files {
		fmt.Println(f.Name())
	}

}
