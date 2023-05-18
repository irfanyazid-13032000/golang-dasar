package main

import "fmt"

type Man struct{
	name string
}

func (man *Man) Married(){
	man.name = "Mr. " + man.name
}

func main(){
	yajed := Man{"yajed"}
	yajed.Married()
	fmt.Println(yajed)
}