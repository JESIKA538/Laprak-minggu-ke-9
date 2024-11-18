package main

import "fmt"

func main() {
	var a int
	var text string
	fmt.Scan(&a)
	text = "negatif"
	if a < 0 && a%2 == 0 {
		text = " Genap negaif"

	} else {
		text = "Bukan"
	}
	fmt.Println(text)
}
