package main

import (
	"fmt"
	"os"
)

func main() {
	readFile("./data/file/sample.txt")
}

func readFile(fp string) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(string(b))
	}
}
