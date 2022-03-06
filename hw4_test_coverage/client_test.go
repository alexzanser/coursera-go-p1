package main

import (
	"net/http"
	"net/http"
	"testing"
	"reflect"
	"io"
)

// type TestCase struct {
// 	ID 			string
// 	Response 	string
// 	StatusCode 	int
// }


type Cart struct {
	PaymentApiURL string
}

func TestCartCheckout(t *testing.T) {
	cases := []TestCase{
		TestCase{
			ID: "42",
			Result: &CheckoutResult {
				Status: 200,
				Balance: 100500,
				Err: "",
			},
			IsError: false,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(CheckoutDummy))

	for caseNum, item := range cases {
		c := &Cart{
			PaymentApiURL: ts.URL,
		}
		result, err := c.Checkout(item.ID)
		if err != nil && !item.IsError {
			t.Errorf("[%d] unexpected error: %#v", caseNum, err)
		}
		if err == nil && item.IsError {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !reflect.DeepEqual(item.Result, result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v",
										caseNum, item.Result, result)
		}
	}
		ts.Close()
}


func CheckoutDummy(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	switch key {
		case "42":
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, `{"status": 200, "balance": 100500}`)
		case "100500":
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, `{"status": 400, "err": "bad_balance"}`)
		case "__broken_json":
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, `{"status": 400`) //broken json
		case "__internal_error":
			fallthrough
		default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
