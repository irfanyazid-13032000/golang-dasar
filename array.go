package main

import "fmt"

func main(){
	var names [2]string

	names[0] = "irfan"
	names[1] = "yajed"

	fmt.Println(names[0])


	var values = [3]int{
		89,
		76,
		87,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	var panjangArray [10]string

	fmt.Println(len(panjangArray))
}