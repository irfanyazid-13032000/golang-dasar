
package main

import "fmt"

func getGoodBye(name string)string{
	return "good bye " + name
}

func main(){
	// function getGoodBye dimasukkan kedalam variable salamPerpisahan
	salamPerpisahan := getGoodBye
	fmt.Println(salamPerpisahan("yajed"))
	
}