package main

import "fmt"

func main() {
	for i := 1; i < 20; i += 2 {
		if i == 20 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Fim")

}
