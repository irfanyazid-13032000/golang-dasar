package main

import "fmt"

func random() interface{}{
	return true
}

func main(){
	var result interface{} = random()
	// result.(string) adalah type assertion untuk mengubah interface kosong ke string
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string loh : ",value)
	case int:
		fmt.Println("integer",value)
	default:
		fmt.Println("haduh type data apa ini",value) 

	}
}