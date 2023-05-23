package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("yajed irpan","yajed"))
	fmt.Println(strings.Split("yajed irpan"," "))
	fmt.Println(strings.ToLower("YAJED IRPAN"))
	fmt.Println(strings.ToTitle("Yajed irpan"))
	fmt.Println(strings.Trim("$Yajed irpan$$$","$"))
	fmt.Println(strings.ReplaceAll("yajed adalah anak yang baik","yajed","tono"))
}