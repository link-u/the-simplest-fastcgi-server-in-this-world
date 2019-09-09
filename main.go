package main

import (
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func handler(w http.ResponseWriter,  req *http.Request) {
	_, _ = w.Write([]byte("zoi!"))
}

func main() {
	var sockFile = os.Args[1]
	http.HandleFunc("/", handler)
	_ = os.Remove(sockFile)
	listener, err := net.Listen("unix", sockFile)
	log.Print("Listen at: ", sockFile)
	if err != nil {
		log.Fatal(err)
	}
	err = fcgi.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
