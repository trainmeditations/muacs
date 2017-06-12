package main

import (
	"fmt"
	"net/http"
//	"golang.org/x/net/http2"
//	"net"
//	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!\n", r.URL.Path[1:])
}

func main() {
	listenAddress := "localhost:8081"
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	s := &http.Server {
		Addr: listenAddress,
		Handler: m,
	}

	//for http2 without tls
	/*s2 := &http2.Server {}
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(err)//TODO Handle
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			//TODO Handle
			log.Print(err)
		} else {
			//s is a net.http.Server and m is a net.http.ServeMux
			//TODO: check for http1.1 upgrade
			go s2.ServeConn(conn, &http2.ServeConnOpts{BaseConfig: s, Handler: m})
		}
	}*/
	
	//for http1.1
	s.ListenAndServe()
}
