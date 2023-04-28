package main

import "fmt"

func main(){
	name := "joko"

	// tidak bisa pake 3 sama dengan ===
	if name == "yajed"{
		fmt.Println("halo yajed")
	}else if name == "jokowi"{
		fmt.Println("halo pak presiden")
	}else{
		fmt.Println("hai, kamu ganteng, tidak seperti yajed")
	}


	// short statement diberikan titik koma
	if length := len(name); length>5{
		fmt.Println("nama terlalu panjang")
	}else{
		fmt.Println("udah pas kok")
	}
}