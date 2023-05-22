package helper

import "fmt"

var version = 1

var Application = "golang"

func SayHello(name string){
	fmt.Println("hello",name)
}

func sayGoodBye(name string){
	fmt.Println("goodbye",name)
}