package main

import "fmt"

func endApp(){
	// Fungsi recover membantu program untuk tetap berjalan dengan menangani kasus panic.
	// recover tidak akan di eksekusi apabila masih dibawah function panic
	// oleh karena itu ada baiknya dimasukkan ke function defer karena function defer akan tetap di eksekusi walaupun ada panic
	message := recover()
	if message != nil{
	fmt.Println("error dgn message:",message)
	}

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
	fmt.Println("yajed")
}