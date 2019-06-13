package httpsServer

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of https service in golang!")
}

func HttpsServer(port string) {
    http.HandleFunc("/", handler)

    fmt.Println("https://localhost:" + port)
    http.ListenAndServeTLS(":" + port, "ssl/cert.crt",
        "ssl/private.key", nil)
}