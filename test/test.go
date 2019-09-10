package main

import (
	"net/http"
	"worst"
)

func main() {
	w := worst.New()


	w.Security.AllowedOrigins = []string{"https://badreputation.pt"}

	w.SetMiddlewareDefaults()
	w.SetSecurityDefaults()

	w.Router.Get("/test", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("asdasdas"))

	})

	w.Run()
}
