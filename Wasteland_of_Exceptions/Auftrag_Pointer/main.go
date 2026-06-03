package main

import "fmt"

func main() {
	x := 42
	
	fmt.Printf("Variable x: %d\n", x)
	fmt.Printf("Address: %p\n", &x)
	
	ptr := &x
	fmt.Printf("Pointer value: %d\n", *ptr)
	
	*ptr = 100
	fmt.Printf("After change: x = %d\n", x)
}
