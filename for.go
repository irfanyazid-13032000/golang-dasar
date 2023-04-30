package main

import "fmt"

func main(){
	
	for counter := 1;counter<= 10;counter++{
		fmt.Println("perulangan ke-",counter)
		
	}

	slice := []string{"irpan","yajed","terbaik"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("index:",i,"value:",value)
		fmt.Println(value)
		
	}

	person := make(map[string]string)
	person["name"] = "yajed"
	person["pekerjaan"] = "pelawak"
	person["visimisi"] = "tidak ada"

	for key,value := range person{
		fmt.Println(key,value)
	}
}