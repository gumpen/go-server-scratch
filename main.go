package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", loggingHandler(handle))
	http.ListenAndServe("127.0.0.1:8080", nil)
}
