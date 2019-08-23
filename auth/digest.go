package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/digest",
		func(w http.ResponseWriter, r *http.Request) {
			if _, ok := r.Header["Authorization"]; !ok {
				w.Header().Add("WWW-Authenticate", `Digest realm="my private area", nonce="RMH1usDrAwA=6dc290ea3304de42a7347e0a94089ff5912ce0de", algorithm=MD5, qop="auth"`)
				w.WriteHeader(http.StatusUnauthorized) // 401
				http.Error(w, "Not authorized", 401)
				return
			}
			_, _ = fmt.Fprintf(w, "Authed!")
		},
	)
	_ = http.ListenAndServe(":18888", nil)
}