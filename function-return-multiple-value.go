package main

import "fmt"

func getFullName() (string,string,int){
	return "irpan","yajed",23
}

func main(){
	firstName,_,_ := getFullName()

	fmt.Println(firstName)
}