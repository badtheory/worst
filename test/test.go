package main

import (
	"net/http"
	"worst"
)

func main() {
	w := worst.New()

	w.SetDefaults()

	w.Router.Get("/test", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("asdasdas"))

	})

	w.Run()
}
