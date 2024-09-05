package main

import (
	"log"
	netHttp "net/http"
	"product-engine/server"
	"product-engine/server/http"
)

func main() {
	err := server.Init()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	defer server.Close()

	router := http.NewServer()

	log.Println("Server starting at :8080")
	log.Fatal(netHttp.ListenAndServe(":8080", router))
}
