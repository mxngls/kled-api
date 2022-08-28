package main

import (
	"log"
	"net/http"

	"github.com/mxngls/kled-server/server"
)

func main() {
	http.HandleFunc("/search", server.Search)
	http.HandleFunc("/view", server.View)

	// Add message to indicate that the server is working
	log.Printf("\rAbout to listen on port 8090. Go to https://127.0.0.1:8090/")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
