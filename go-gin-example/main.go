package main

import (
	"fmt"
	"github.com/anhoder/go-gin-example/pkg/setting"
	"github.com/anhoder/go-gin-example/routers"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}