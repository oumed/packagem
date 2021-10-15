package main

import (
	"fmt"
	a "github.com/oumed/packagea"
	b "github.com/oumed/packageb"
)

func main(){
	//fmt.Println("Hi every one ...")
	aa := a.GetName("HHHHHHH")
	bb := b.GetMessage()
	fmt.Println(aa)
	fmt.Println(bb)
}
