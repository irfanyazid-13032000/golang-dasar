package main

import "fmt"

func main(){
	
	name := "yajed"
	counter := 0

	increment := func ()  {
		// membuat variable baru dengan nama name (berbeda dengan yang diatas)
		name := "budi"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(name)


}