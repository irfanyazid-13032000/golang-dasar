package main

import "fmt"

func sayHelloWithFilter(name string,filter func(string) string){
	// filter(name) adalah untuk menjalankan method filterKalimatKotor() yg sudah dikirim melalui parameter
	nameFiltered := filter(name)
	fmt.Println("hello",nameFiltered)
}

func filterKalimatKotor(name string)string{
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main(){
	sayHelloWithFilter("anjing",filterKalimatKotor)
}