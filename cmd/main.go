package main

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"

	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/fib", GetFibonacci),
	)
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetFibonacci(w rest.ResponseWriter, r *rest.Request) {
	number := r.PathParam("fib")

	fmt.Println(number)
	w.WriteJson(map[string]string{"result": number})
}
