package main

import "fmt"

func logging(){
	fmt.Println("selesai memanggil logging")
}

func runApplication(value int){

	defer logging()

	fmt.Println("run application...")


	hasil := 10/value

	fmt.Println("hasilnya adalah",hasil)

	
}

func main(){
	runApplication(0)
}