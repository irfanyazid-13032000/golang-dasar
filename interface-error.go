package main

import (
	"fmt"
	"errors"
)

func Pembagian(nilai int,pembagi int) (int,error){
	if pembagi == 0{
		return 0,errors.New("haduh gak bisa dong dibagi nol")
	}else{
		return nilai/pembagi,nil
	}
}

func main(){
	// var contohError error = errors.New("haduh error nih")

	cobaBagiAh,err := Pembagian(200,0)

	if err == nil {
		fmt.Println("hasil",cobaBagiAh)
	}else{
		fmt.Println("error",err.Error())
	}

}