package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse is a wrapper for all JSON responses
type JSONResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

// JSONHandler is an HTTP handler tailored for JSON responses, for use with the ServeJSON HTTP handler.
type JSONHandler interface {
	ServeJSON(http.ResponseWriter, *http.Request) (interface{}, int, error)
}

// JSONHandlerFunc converts a function with signature `func(http.ResponseWriter, *http.Request) (interface{}, int, error)` to a JSONHandler
type JSONHandlerFunc func(http.ResponseWriter, *http.Request) (interface{}, int, error)

// ServeJSON calls f(w, r).J
func (f JSONHandlerFunc) ServeJSON(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	return f(w, r)
}

// ServeJSON is an HTTP handler that allows other HTTP handlers to return data and a status code for JSON serialization. Also provides error handling.
func ServeJSON(handler JSONHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		var response JSONResponse

		if data, status, err := handler.ServeJSON(w, r); err != nil {
			if status == 0 {
				status = http.StatusInternalServerError
			}

			response = JSONResponse{
				Data:   nil,
				Error:  fmt.Sprint(err),
				Status: false,
			}

			SendJsonResponse(w, status, response)
		} else {
			response = JSONResponse{
				Data:   data,
				Error:  nil,
				Status: true,
			}

			SendJsonResponse(w, status, response)
		}
	})
}

func SendJsonResponse(w http.ResponseWriter, status int, response JSONResponse) {
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(response)
}
