package main

import "fmt"

func main() {
	for i := 0; i < 31; i += 3 {
		if i == 31 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("fim")
}
