package main

import (
	"fmt"
	"github.com/jiangboLee/go_web_study/leechitchat/data"
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

func postThread(rw http.ResponseWriter, req *http.Request) {
	sess, err := session(rw, req)
	if err != nil {
		http.Redirect(rw, req, "/login", 302)
	} else {
		err = req.ParseForm()
		if err != nil {
			danger(err, "不能格式表单")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "不能根据session找到用户")
		}
		body := req.PostFormValue("body")
		uuid := req.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(rw, req, "读取不到这个帖子")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "不能回复")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(rw, req, url, 302)
	}
}

func readThread(rw http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(rw, req, "读取不到这个帖子")
	} else {
		_, err := session(rw, req)
		if err != nil {
			generateHTML(rw, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(rw, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}
