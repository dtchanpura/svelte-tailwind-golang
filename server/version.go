package server

import "net/http"

const (
	// Version specifies the active version
	Version = "latest"
)

func handleVersion() func(http.ResponseWriter, *http.Request) {
	type v struct {
		Version string `json:"version"`
	}
	return handleResponse(Response{Ok: true, Result: v{Version: Version}, StatusCode: http.StatusOK})
}
