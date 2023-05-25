package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())


	utc := time.Date(2020,12,25,9,40,20,34,time.UTC)

	fmt.Println(utc)

	layout := "2006-01-02"
	parse,_ := time.Parse(layout,"2020-12-01")
	fmt.Println(parse)
}