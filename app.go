package main

import (
 "fmt"
 "log"
 "net/http"
 "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Content-Type", "text/plain; charset=utf-8")
 fmt.Fprintln(w, "Hello")
 fmt.Fprintf(w, "Время сервера: %s\n", time.Now().Format("2006-01-02 15:04:05"))
 fmt.Fprintln(w, "Контрольная работа: Вариант 11")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
 w.WriteHeader(http.StatusOK)
 fmt.Fprintln(w, "OK")
}

func main() {
 http.HandleFunc("/", handler)
 http.HandleFunc("/health", healthHandler)
 log.Fatal(http.ListenAndServe(":8888", nil))
}
