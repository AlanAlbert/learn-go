package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func Index(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(writer, "Hello %s!", ps.ByName("name"))
}

func Time(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	now := time.Now()
	timeStr := fmt.Sprintf("{\"timestamp\": %d}", now.Unix())
	_, _ = fmt.Fprintln(writer, timeStr)
}

func main()  {
	router := httprouter.New()
	router.GET("/hello/:name", Index)
	router.GET("/time", Time)

	_ = http.ListenAndServe(":8080", router)
}


