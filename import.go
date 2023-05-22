package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main(){
	helper.SayHello("yajed")
	// helper.sayGoodBye("yajed") //error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error
}