package main

import "fmt"

func endApp(){
	fmt.Println("aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("APLIKASI ERROR LOHHHH")
	}
	// apabila panic dieksekusi, maka fungsi bawahnya tidak akan dieksekusi
	fmt.Println("aplikasi berjalan")
}

func main(){
	runApp(true)
}