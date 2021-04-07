package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "",""
	// os.Arg[0]表示程序的名字,而实际的参数在 1:end里面存储
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}