package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetParameter(t *testing.T) {
	tests := []struct {
		name    string
		req     *http.Request
		want    int
		wantErr bool
	}{{
		name:    "get parameter from http request",
		req:     httptest.NewRequest("GET", "/fib?n=0", nil),
		want:    0,
		wantErr: false,
	}, {
		name:    "get parameter of not allowed data type from http request",
		req:     httptest.NewRequest("GET", "/fib?n=a", nil),
		want:    0,
		wantErr: true,
	}, {
		name:    "get parameters from http request is not possible",
		req:     httptest.NewRequest("GET", "/fib?n=", nil),
		want:    0,
		wantErr: true,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetParameter(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParameter() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name    string
		number  int
		want    string
		wantErr bool
	}{{
		name:    "calculate Fibonacci Numbers",
		number:  1,
		want:    "1",
		wantErr: false,
	}, {
		name:    "calculate Fibonacci Numbers with negative numbers",
		number:  -1,
		want:    "<nil>",
		wantErr: true,
	}, {
		name:    "calculate Fibonacci Numbers with large numbers",
		number:  100,
		want:    "354224848179261915075",
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fibonacci(tt.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})

	}
}
