package main

import "fmt"

type Address struct{
	city,province,country string
}

func ChangeAddresToIndonesia(address *Address){
	address.country = "indonesia"
}

func main(){

	address1 := Address{"mojosongo","jawa timur","kanada"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.city = "sumobito"

	*address2 = Address{"bekasi","jawa barat","kanada"}


	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.city = "jombang"

	fmt.Println(address4)

	var alamat = Address{
		city: "cikampek",
		province: "jawa barat",
		country: "malaysia",
	}

	ChangeAddresToIndonesia(&alamat)

	fmt.Println(alamat)
}