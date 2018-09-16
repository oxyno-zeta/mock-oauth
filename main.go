package main

import (
    "flag"
    "log"
    "net/url"
    "net/http"
)

var (
    message string
    address string
)


func handleAuth(w http.ResponseWriter, r *http.Request) {


	whereToRedirect := r.Referer()
	if whereToRedirect == "" {
		whereToRedirect = "http://example.org?no=1234"
	}
	redirectURL, _ := url.Parse(whereToRedirect)
	redirectURL.Query().Set("state", r.URL.Query().Get("state"))
	redirectURL.Query().Set("code", "randomcode")

	w.Header().Set("Location", redirectURL.String())
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleToken(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"token\":\"sometoken\"}"))
}



func handleUserInfo(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("{\"name\":\"johndoe\",\"email\":\"a@a.com\",\"username\":\"someusername\"}"))
}


func init() {
    flag.StringVar(&message, "message", "there", "a message to print")
    flag.StringVar(&address, "address", "0.0.0.0:8080", "address/port to listen on")
}
func main() {
    flag.Parse()
    http.HandleFunc("/auth", handleAuth)
	http.HandleFunc("/token", handleToken)
	http.HandleFunc("/userinfo",handleUserInfo)
	/*
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", message)
    })
    */
    log.Printf("listening on %s...", address)
    log.Fatal(http.ListenAndServe(address, nil))
}