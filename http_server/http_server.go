package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "Hello world!")
	})
	http.HandleFunc("/time/", func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		timeStr := fmt.Sprintf("{\"timestamp\": %d}", now.Unix())
		_, _ = fmt.Fprintln(writer, timeStr)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
