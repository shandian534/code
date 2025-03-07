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
	//文件读取操作
	contents, err := os.ReadFile("/Users/shandian/GolandProjects/code/chapter3/wordcount/gowords.txt")
	//错误处理
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}
	//数据类型转换
	text := string(contents)
	//调用函数
	count := words.CountWords(text)
	//输出
	fmt.Printf("There are %d words in your text. \n", count)
}
