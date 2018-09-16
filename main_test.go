package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
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

    
}