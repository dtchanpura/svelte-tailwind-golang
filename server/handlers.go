package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct
type Response struct {
	Ok         bool        `json:"ok"`
	Result     interface{} `json:"result,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"-"`
}

func handleResponse(resp Response) func(http.ResponseWriter, *http.Request) {
	if resp.StatusCode == 0 {
		resp.StatusCode = http.StatusOK
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("error encoding: %+v", resp)
			log.Println(err)
			resp.StatusCode = http.StatusInternalServerError
		}
	}
}
