package main

import (
	"github.com/jiangboLee/go_web_study/leechitchat/data"
	"net/http"
)

func index(rw http.ResponseWriter, req *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(rw, req, "Cannot get threads")
	} else {
		_, err := session(rw, req)
		if err != nil {
			P("没有session")
			generateHTML(rw, threads, "layout", "public.navbar", "index")
		} else {
			P("有session")
			generateHTML(rw, threads, "layout", "private.navbar", "index")
		}
	}
}

// GET /err?msg=
// shows the error message page
func err(rw http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	_, err := session(rw, req)
	if err != nil {
		generateHTML(rw, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(rw, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
