package main

import "fmt" 

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		a = a/2
		
	} else {
		a = (a/2 + 1)
	}
	fmt.Println(a)
}