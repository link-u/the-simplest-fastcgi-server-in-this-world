package main

import (
	"fmt"
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
	fmt.Printf("Listen at: %s\n", sockFile)
	listener, err := net.Listen("unix", sockFile)
	if err != nil {
		panic(err)
	}
	err = fcgi.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}
