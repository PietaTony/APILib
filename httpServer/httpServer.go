package httpServer

import (
	"fmt"
	"net/http"
)

func HttpServer(port string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":" + port, nil)

	fmt.Println("http server is running at http://localhost:" + port)
}