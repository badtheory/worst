package main

import (
	"net/http"
	"worst"
)

func main() {
	w := worst.New(worst.Options{
		Server: &http.Server{
			Addr:"localhost:1338",
		},
		Static: worst.Static{
			Url:"/*",
			Path: "/home/paulo/Desktop/public",
		},
	})

	w.Run()

}