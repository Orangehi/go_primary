package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	th := http.RedirectHandler("http://example.org", 307)

	mux.Handle("/login", th)

	http.ListenAndServe(":3000", mux)

}
