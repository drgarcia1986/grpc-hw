package main

import (
	"net"
	"net/http"
)

func httpServe(l net.Listener) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}
