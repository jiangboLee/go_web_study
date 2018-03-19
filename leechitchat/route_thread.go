package main

import (
	_ "./data"
	// "fmt"
	"net/http"
)

func newThread(rw http.ResponseWriter, req *http.Request) {
	_, err := session(rw, req)
	if err != nil {
		http.Redirect(rw, req, "/login", 302)
	} else {
		generateHTML(rw, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(rw http.ResponseWriter, req *http.Request) {
	sess, err := session(rw, req)
	if err != nil {
		http.Redirect(rw, req, "/login", 302)
	} else {
		err = req.ParseForm()
		if err != nil {
			danger(err, "不能格式表单")
		}
		P(sess)
		user, err := sess.User()
		if err != nil {
			danger(err, "不能从session中等到user")
		}
		topic := req.PostFormValue("topic")
		P(topic)
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "不能创建帖子")
		}
		http.Redirect(rw, req, "/", 302)
	}

}
