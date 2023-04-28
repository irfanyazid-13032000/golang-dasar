package main

import "fmt"

func main(){
	person := map[string]string{
		"name":"yajed",
		"address":"mojokerto",
	}

	person["profesi"] = "pelawak"

	fmt.Println(person)


	var book map[string]string = make(map[string]string)

	book["author"] = "gege akutami"
	book["publisher"] = "shounen jump"
	book["title"] = "jujutsu kaisen"
	book["salah"] = "aduh salah nih"
	fmt.Println(book)
	fmt.Println(len(book))


	delete(book,"salah")

	fmt.Println(book)
	fmt.Println(len(book))
}