package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	fmt.Println("Wasm init")
	// Update()
	<-c
}
