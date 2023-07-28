package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	fmt.Println("Wasm init")
	// Update()
	<-c
}
