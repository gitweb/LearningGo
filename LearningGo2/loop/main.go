package main

import "fmt"

func main() {
	for i := 0; i < 2000; i++ {
		fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	}
}
