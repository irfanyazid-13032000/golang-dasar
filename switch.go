package main

import "fmt"

func main(){
	name := "doniki"

	switch name {
	case "yajed":
		fmt.Println("halo yajed")
	case "jokowi":
		fmt.Println("halo pak presiden")
	default:
		fmt.Println("wah rupanya aku tidak mengenalmu")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama pas kok")
	// }


	length:= len(name)
	switch {
	case length>10:
		fmt.Println("nama terlalu panjang")
	case length>5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("sudah benar")
	}
}