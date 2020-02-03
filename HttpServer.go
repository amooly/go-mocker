package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	HttpConfig *Config
}

func (server HttpServer) Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, route := range server.HttpConfig.Routes {
			if route.Request.Url == r.URL.Path {
				if route.Request.Method == r.Method {
					w.WriteHeader(route.Response.Status)
					if _, err := w.Write([]byte(route.Response.Body)); err != nil {
						log.Println("error: failed to write response:", err)
					}
					return
				}
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed))); err != nil {
			log.Println("error: failed to write response:", err)
		}
	})

	addr := fmt.Sprintf("%s:%d", server.HttpConfig.Server.Address, server.HttpConfig.Server.Port)
	log.Println("start to listen:", addr)

	return http.ListenAndServe(addr, nil)
}
