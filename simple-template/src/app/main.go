package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	h := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "8080"
	}

	addr := ":" + port

	l := log.New(os.Stdout, "[App]", log.LstdFlags)
	l.Printf("Start listening on %s...", addr)

	http.ListenAndServe(addr, http.HandlerFunc(h))
}
