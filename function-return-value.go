package main

import "fmt"

func getHello(name string) string{
	return "hello "+name
}

func main(){
	hello := getHello("yazid")
	fmt.Println(hello)
}