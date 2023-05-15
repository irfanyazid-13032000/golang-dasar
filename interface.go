package main

import "fmt"

type HasName interface{
	GetName() string
}

func sayHalo(hasName HasName){
	fmt.Println("hello",hasName.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

type Anime struct{
	Name string
}

func (anime Anime) GetName() string{
	return anime.Name
}

func main(){

	var yajed Person
	yajed.Name = "eko"

	sayHalo(yajed)

	gojo := Anime{
		Name: "gojo",
	}

	sayHalo(gojo)
}
