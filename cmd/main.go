package main

import (
	"fmt"
	"log"

	"github.com/ant0ine/go-json-rest/rest"

	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/fib" {
			fmt.Println(r.URL.Query())
			w.Write([]byte(r.URL.Query().Get("n")))
		}
	}
}

func GetFibonacci(w rest.ResponseWriter, r *rest.Request) {
	number := r.PathParam("fib")

	fmt.Println(number)
	w.WriteJson(map[string]string{"result": number})
}
