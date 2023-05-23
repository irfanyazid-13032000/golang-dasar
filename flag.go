package main

import (
	"flag"
	"fmt"
)

func main() {
	var nama *string = flag.String("nama","yajid","masukkin nama")

	flag.Parse()

	fmt.Println(*nama)
}