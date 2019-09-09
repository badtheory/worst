package main

import (
	"worst"
	"net/http"
	"time"
)

func main() {
	w := worst.New(worst.Options{
		Server: &http.Server{
			Addr:"localhost:1341",
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	})
	w.Run()
}
