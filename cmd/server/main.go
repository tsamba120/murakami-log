package main

import (
	"log"

	"github.com/tsamba120/murakami-log/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
