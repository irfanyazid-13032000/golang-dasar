package main

import (
	"fmt"
	"os"
)

func main(){
	
	fmt.Println(os.Args)

	name,err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	}else{
		fmt.Println("error nih")
	}

	username := os.Getenv("APP_USERNAME")

	fmt.Println(username)
}