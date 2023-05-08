package main

import "fmt"

// continue akan skip perintah pada perulangan
func main(){
	for i := 0; i < 10; i++ {
		if i%2 == 1 && i>0 {
			continue
		}
		fmt.Println("perulangan ke",i)
		
	}
}