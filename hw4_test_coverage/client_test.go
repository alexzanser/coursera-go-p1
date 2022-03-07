package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	// "reflect"
	// "fmt"
	"io"
)

const (
	secretToken = "vasya228"
)

type TestCase struct {
	Request		SearchRequest
	Result		*SearchResponse
	IsError 	bool
	Token 		string
}

// type User struct {
// 	Id     int
// 	Name   string
// 	Age    int
// 	About  string
// 	Gender string
// }

// type SearchRequest struct {
// 	Limit      int
// 	Offset     int    // Можно учесть после сортировки
// 	Query      string // подстрока в 1 из полей
// 	OrderField string
// 	// -1 по убыванию, 0 как встретилось, 1 по возрастанию
// 	OrderBy int
// }

// type SearchResponse struct {
// 	Users    []User
// 	NextPage bool
// }

// type Cart struct {
// 	PaymentApiURL string
// }

func SearchServer(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("limit")
	if r.Header.Get("AccessToken") != secretToken {
		w.WriteHeader(http.StatusUnauthorized)
	} else if r.FormValue("query") == "invalid_json" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{}}`)
	} else if r.FormValue("order_field") != "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"Error" : "ErrorBadOrderField"}`)
	} else if r.FormValue("query") == "unknown_request_error" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{}`)
	} else if r.FormValue("query") == "server_fatal" {
		w.WriteHeader(http.StatusInternalServerError)
	} else if r.FormValue("query") == "unknown_error" {
		w.WriteHeader(-1)
	} else if r.FormValue("query") == "timeout" {
		w.WriteHeader(http.StatusProcessing  )
	} else if r.FormValue("query") == "wrong_result_json" {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{}}`)
	} else {
		switch key {
			case "2":
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, `[{"Id": 1, "Name": "Boyd Wolf", "Age": 22, "About": "gay", "Gender":"male"}]`)
			case "3":
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, `[{"Id": 1, "Name": "Boyd Wolf", "Age": 22, "About": "gay", "Gender":"male"}, 
				{"Id": 1, "Name": "Boyd Wolf", "Age": 23, "About": "gay", "Gender":"male"},
				{"Id": 1, "Name": "Boyd Wolf", "Age": 24, "About": "gay", "Gender":"male"}]`)

			// case "__broken_json":
			// 	w.WriteHeader(http.StatusOK)
			// 	io.WriteString(w, `{"status": 400`) //broken json
			case "__internal_error":
				fallthrough
			default:
				w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func TestCartCheckout(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "",
				OrderField: "Z",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: "wrongToken",
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "invalid_json",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "unknown_request_error",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "server_fatal",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 4,
				Offset: 1,
				Query: "wrong_result_json",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 0},
				},
				NextPage: false,
			},
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 1,
				Offset: 0,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 1, Name: "Boyd Wolf", Age: 22, About: "gay", Gender:"male"},
				},
				NextPage: false,
			},
			IsError: false,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 2,
				Offset: 0,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: &SearchResponse {
				Users: []User{
					{Id: 1, Name: "Boyd Wolf", Age: 22, About: "gay", Gender:"male"},
					{Id: 1, Name: "Boyd Wolf", Age: 23, About: "gay", Gender:"male"},
				},
				NextPage: true,
			},
			IsError: false,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: -1,
				Offset: 0,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: nil,
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 30,
				Offset: 0,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: nil,
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 1,
				Offset: -1,
				Query: "",
				OrderField: "",
				OrderBy: 0,
			},
			Result: nil,
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 1,
				Offset: 1,
				Query: "unknown_error",
				OrderField: "",
				OrderBy: 0,
			},
			Result: nil,
			IsError: true,
			Token: secretToken,
		},
		TestCase{
			Request:  SearchRequest {
				Limit: 1,
				Offset: 1,
				Query: "timeout",
				OrderField: "",
				OrderBy: 0,
			},
			Result: nil,
			IsError: true,
			Token: secretToken,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	for caseNum, item := range cases {
		c := &SearchClient{
			URL: ts.URL,
			AccessToken: item.Token,
		}
		_, err := c.FindUsers(item.Request)
		if err != nil && !item.IsError {
			t.Errorf("[%d] unexpected error: %#v", caseNum, err)
		}
		// if err != nil && item.IsError {
		// 	t.Errorf("[%d] expected error '%s', got nil", caseNum, err)
		// }
		// if !reflect.DeepEqual(item.Result, result) {
		// 	t.Errorf("[%d] wrong result, got %#v",
		// 								caseNum, result)
		// }
	}
	ts.Close()
}
