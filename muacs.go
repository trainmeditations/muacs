package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hkwi/h2c"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!\n", r.URL.Path[1:])
}

func main() {
	listenAddress := "localhost:8081"
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	s := &http.Server{
		Addr:    listenAddress,
		Handler: &h2c.Server{Handler: m, DisableDirect: false},
	}
	log.Fatal(s.ListenAndServe())
}
