package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("yajed")
	data.PushBack("ganteng")
	data.PushBack("banget")
	data.PushFront("tono")



	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}