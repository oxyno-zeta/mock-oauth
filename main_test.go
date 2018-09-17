package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth", handleAuth)
	return httptest.NewServer(mux)
}

func TestHandleAuth(t *testing.T) {
	server := createServer()
	idpServerURL := "http://" + server.Listener.Addr().String()
	fmt.Println(idpServerURL)

	// TODO: mount handlers in a new server and test it.

}
