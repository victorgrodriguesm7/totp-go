package main

import (
	"fmt"
	"net/http"
)

func totpRedirectHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email");
	secret := r.URL.Query().Get("secret");
	issuer := r.URL.Query().Get("issuer");

	if email == "" {
		email = "example@example.com";
	}

	if secret == "" {
		secret = "EXAMPLEOFSECRET1";
	}

	if issuer == "" {
		issuer = "default";
	}
	

	url := fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", email, secret, issuer);

	fmt.Print(url);

	http.Redirect(w, r, url, http.StatusFound);
}


func main() {

    http.HandleFunc("/", totpRedirectHandler)

    http.ListenAndServe(":8080", nil)
}