package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"strconv"

	"net/http"
)

type Response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/fib" {
			n, err := GetParameter(r)
			if err != nil {
				response := CreateErrorResponse(err, http.StatusBadRequest)
				json.NewEncoder(w).Encode(&response)
				return
			}

			number, err := Fibonacci(n)
			if err != nil {
				response := CreateErrorResponse(err, http.StatusInternalServerError)
				json.NewEncoder(w).Encode(&response)
				return
			}

			response := Response{
				Result: number.String(),
			}
			json.NewEncoder(w).Encode(&response)
		} else {
			err := errors.New("unknown parameters")
			response := CreateErrorResponse(err, http.StatusBadRequest)
			json.NewEncoder(w).Encode(&response)
		}
	default:
		err := errors.New("this method is not allowed")
		response := CreateErrorResponse(err, http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(&response)
	}
}

func CreateErrorResponse(err error, status int) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Message: err.Error(),
	}
}

func GetParameter(r *http.Request) (int, error) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func Fibonacci(n int) (*big.Int, error) {
	if n < 0 {
		err := errors.New("not supported for values less than 0")
		return nil, err
	}
	// TODO:nの上限設定

	x, y := big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		x, y = y, new(big.Int).Add(x, y)
	}
	return x, nil

}
