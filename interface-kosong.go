package main

import "fmt"

func Ups(i int) interface{}{
	if i == 1 {
		return 1
	}else if i == 2 {
		return 2
	}else{
		return true
	}
}

func main(){
	var internih interface{} = Ups(1)
	fmt.Println(internih)
}