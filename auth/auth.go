// basic認証
package main

import (
"fmt"
"net/http"
)

// user情報のダミーデータ
const (
basicAuthUser     = "user"
basicAuthPassword = "pass"
)

// Basic認証
func main() {
http.HandleFunc("/basic",
func(w http.ResponseWriter, r *http.Request) {
if user, pass, ok := r.BasicAuth(); !ok || user != basicAuthUser || pass != basicAuthPassword {
w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
w.WriteHeader(http.StatusUnauthorized) // 401
http.Error(w, "Not authorized", 401)
return
}
// 認証されれば"Authed!!"を返す.
	_, _ = fmt.Fprintf(w, "Authed!!")
},
)
	_ = http.ListenAndServe(":18888", nil)
}