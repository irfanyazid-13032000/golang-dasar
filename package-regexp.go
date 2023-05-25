package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`y([a-z])d`)
	// dari youtube programmer zaman now
	var youtube *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("yad"))

	search := youtube.FindAllString("eko edo eyo ela edy",-1)
	fmt.Println(search)
}