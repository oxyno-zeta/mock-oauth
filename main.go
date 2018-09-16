package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

var (
	message string
	address string
)

func handleAuth(w http.ResponseWriter, r *http.Request) {

	whereToRedirect := r.URL.Query().Get("redirect_uri")
	if whereToRedirect == "" {
		whereToRedirect = "http://example.org?no=1234"
	}
	redirectURL, _ := url.Parse(whereToRedirect)
	params := redirectURL.Query()
	params.Set("state", r.URL.Query().Get("state"))
	params.Set("code", "randomcode")

	redirectURL.RawQuery = params.Encode()

	w.Header().Set("Location", redirectURL.String())
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"token\":\"sometoken\"}"))
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"name\":\"johndoe\",\"email\":\"a@a.com\",\"preferred_username\":\"shbose\",\"sub\":\"62ccaef-5bd6-48c8-bbb1-987d39482a02\"}"))
}

func init() {
	flag.StringVar(&message, "message", "there", "a message to print")
	flag.StringVar(&address, "address", "0.0.0.0:8080", "address/port to listen on")
}
func main() {
	flag.Parse()
	http.HandleFunc("/auth", handleAuth)
	http.HandleFunc("/token", handleToken)
	http.HandleFunc("/userinfo", handleUserInfo)
	/*
		    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello %s\n", message)
		    })
	*/
	log.Printf("listening on %s...", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
