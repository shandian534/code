// Sample program to show how to show you how to briefly work with io.
package main

import (
	"fmt"
	"github.com/goinaction/code/chapter3/words"
	"os"
)

// main is the entry point for the application.
func main() {
	//filename := os.Args[1]

	contents, err := os.ReadFile("/Users/shandian/GolandProjects/code/chapter3/wordcount/gowords.txt")
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
