package main

import (
	"flag"
	"fmt"
)

func main() {
	surname := flag.String("surname","wang", "姓氏")
	var personalName string
	flag.StringVar(&personalName,"personalName","小二","名字")
	id := flag.Int("id",0,"id")
	flag.Parse()
	fmt.Printf("I am %v %v and my id is %v\n", *surname, personalName,*id)
}