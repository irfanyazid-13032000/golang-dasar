package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("haduh")
	if err == nil{
		fmt.Println(boolean)
	}else{
		fmt.Println(err.Error())
	}

	number,err := strconv.ParseInt("haduh salah nih",2,64)
	if err == nil{
		fmt.Println(number)
	}else{
		fmt.Println(err)
	}

	value := strconv.FormatInt(1000000,16)
	fmt.Println(value)

	intvalue,_ := strconv.Atoi("2000000")


	fmt.Println(intvalue)

}