package main

import "fmt"

func sumAll(numbers ...int)int {
	total := 0

	for _,value := range numbers{
		total += value
	}

	return total
}

func main(){
	fmt.Println(sumAll(123,123,1242,234,1423,5,25))

	// apabila data yg dimiliki adalah slice, maka berikan ... dibelakang nya
	slice := []int{12,34,13,423,534,234,13}
	fmt.Println(sumAll(slice...))
}