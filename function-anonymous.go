package main

import "fmt"

// type ini adalah untuk menulis di parameter, gak ada hubungannya dengan function asli nya
type Blacklist func(string)bool

func registerUser(name string,blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("kamu ditolak",name)
	}else{
		fmt.Println("selamat datang",name)
	}
}

func main(){
	// blacklist merupakan function aslinya
	blacklist := func(name string)bool{
		return name == "anjing"
	}

	// mengirim function blacklist sebagai parameter
	registerUser("doni",blacklist)

	// mengirim function dengan membuat nya langsung
	registerUser("anjing",func(name string) bool {
		return name == "anjing"
	})
}