package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	if err := os.Chdir(".."); err != nil {
		log.Fatal(err)
	}

	srv, err := server.InitServer(log.Default())

	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(srv.Server.Addr, srv.Server.Handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
