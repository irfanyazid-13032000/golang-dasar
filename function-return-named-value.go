package main

import "fmt"

func getPenuhName() (firstName string,lastName string){
	firstName = "irfan"
	lastName = "yajed"

	return
}

func main(){

	firstName,lastName := getPenuhName()

	fmt.Println(firstName,lastName)
	
}