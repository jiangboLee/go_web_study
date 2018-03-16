package main

import (
	"net/http"
	"time"
)

func main() {
	p("LeeChitchat", version(), "started at", config.Address)

	mux := http.NewServeMux()
	//静态文件处理
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
