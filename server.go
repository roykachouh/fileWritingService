package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	f, _ := os.Create("/tmp/log.txt")

	defer f.Close()

	log.SetOutput(f)

	log.Println("Listening on 8081")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
