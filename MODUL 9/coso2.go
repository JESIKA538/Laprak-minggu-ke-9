package main

import "fmt"

func main() {
	var a int
	var text string
	fmt.Scan(&a)
	text = "negatif"
	if a > 0 {
		text = "positif"

	}
	fmt.Println(text)
}
