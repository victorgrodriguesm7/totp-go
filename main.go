package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var indexHTML []byte

func removeFirstSlash(s string) string {
	if strings.HasPrefix(s, "/") {
		return s[1:]
	}
	return s
}

func totpRedirectHandler(w http.ResponseWriter, r *http.Request) {
	data, err := base64.StdEncoding.DecodeString(removeFirstSlash(r.URL.Path));

	if err != nil {
		http.Error(w, "Invalid Base64", http.StatusBadRequest)
		return
	}

	parts := strings.Split(string(data), ":")

	if len(parts) != 3 {
		http.Error(w, "Invalid Format. Expected label:issuer:secret", http.StatusBadRequest)
		return
	}

	email, secret, issuer := parts[0], parts[1], parts[2]

	if email == "" {
		email = "example@example.com";
	}

	if secret == "" {
		secret = "EXAMPLEOFSECRET1";
	}

	if issuer == "" {
		issuer = "default";
	}
	

	url := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, email, secret, issuer);

	log.Println(url);

	http.Redirect(w, r, url, http.StatusFound);
}

func indexHandler(w http.ResponseWriter, _ *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Write(indexHTML)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		indexHandler(w, r);

		return;
	}

	totpRedirectHandler(w, r);
}

func main() {
	var err error;
	log.Println("Reading index.html");

	indexHTML, err = os.ReadFile("assets/index.html")

	if err != nil {
		log.Fatalf("Erro ao carregar index.html: %v", err)
	}

	log.Println("Starting server");

    http.HandleFunc("/", httpHandler);

    log.Fatal(http.ListenAndServe(":8080", nil));
}