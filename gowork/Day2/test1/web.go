package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func helloserver(w http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("h1.gtpl")

	log.Println(t.Execute(w, nil))

}

func h1(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("h2.gtpl")
	log.Println(t.Execute(w, nil))

}

func main() {

	http.HandleFunc("/", helloserver)

	http.HandleFunc("/h2", h1)

	err := http.ListenAndServe("127.0.0.1:9090", nil)

	if err != nil {

		fmt.Println(err)

	}

}
