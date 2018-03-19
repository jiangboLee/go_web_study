package main

import (
	"net/http"
	"time"
)

func main() {
	P("LeeChitchat", Version(), "started at", Config.Address)

	mux := http.NewServeMux()
	//静态文件处理
	files := http.FileServer(http.Dir(Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//首页
	mux.HandleFunc("/", index)
	//error
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signup_account)
	mux.HandleFunc("/logout", logout)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)

	server := &http.Server{
		Addr:           Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
