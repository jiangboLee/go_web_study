package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello word")
}

func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("method:  ", req.Method)
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(rw, nil))
	} else {
		req.ParseForm()
		fmt.Println("username: ", req.Form["username"])
		fmt.Println("password: ", req.Form["password"])
		fmt.Fprintf(rw, "登录成功")
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
