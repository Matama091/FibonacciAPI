package main

import (
	"fmt"
	"log"
	"strconv"

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
			n, err := strconv.Atoi(r.URL.Query().Get("n"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			number := Fibonacci(n)
			fmt.Println(number)

			w.Write([]byte(strconv.Itoa(number)))
		}
	}
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
