package main

import "fmt"

func main() {
	for i := 0; i < 21; i += 2 {
		if i == 21 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("fim com break")
}
