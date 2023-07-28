package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
	dir  = "D:\\particleLife\\assets"
)

func main() {
	log.Printf("lisetnini on %s port, and serve %s directory", port, dir)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}
