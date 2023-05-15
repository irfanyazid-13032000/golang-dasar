package main

import "fmt"

type Customer struct{
	Name,Addres string
	Age int
}

func (pelanggan Customer) saySelamatDatang(name string){
	fmt.Println("selamat datang wahai",pelanggan.Name, "saya adalah",name)
}

func main(){

	var Yajed Customer
	Yajed.Name = "Yajed"
	Yajed.Addres = "Mojosongo"
	Yajed.Age = 23

	Yajed.saySelamatDatang("karyawan")

	fmt.Println(Yajed)
	fmt.Println(Yajed.Name)
	fmt.Println(Yajed.Addres)
	fmt.Println(Yajed.Age)

	// cara lain mengisi struct dengan data

	yono := Customer{
		Name: "yono",
		Addres: "indonesia",
		Age: 67,
	}

	fmt.Println(yono)


	// cara lain mengisi struct (tidak disarankan)

	buno := Customer{"buno","jakarta",25}

	fmt.Println(buno)
	
}