package main

import (
	"github.com/jiangboLee/go_web_study/leechitchat/data"
	"net/http"
)

func login(rw http.ResponseWriter, req *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(rw, nil)
}

//用户提交
func authenticate(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	user, err := data.UserByEmail(req.PostFormValue("email"))
	if err != nil {
		danger(err, "不能找到用户")
	}
	if user.Password == data.Encrypt(req.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "不能创建session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		P("设置cookie")
		P(session)
		http.SetCookie(rw, &cookie)
		http.Redirect(rw, req, "/", 302)
	} else {
		http.Redirect(rw, req, "/login", 302)
	}
}

//用户注册
func signup(rw http.ResponseWriter, req *http.Request) {
	generateHTML(rw, nil, "login.layout", "public.navbar", "signup")
}

func signup_account(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		danger(err, "不能分析表单")
	}
	user := data.User{
		Name:     req.PostFormValue("name"),
		Email:    req.PostFormValue("email"),
		Password: req.PostFormValue("password"),
	}
	P("用户")
	P(user)
	err = user.Create()
	if err != nil {
		danger(err, "不能创建用户")
	}
	http.Redirect(rw, req, "/login", 302)
}

//用户退出
func logout(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(rw, req, "/", 302)
}
