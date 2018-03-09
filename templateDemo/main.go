package main

import (
	"html/template"
	"net/http"
	// "os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func hello(rw http.ResponseWriter, req *http.Request) {
	f1 := Friend{"minux.ma"}
	f2 := Friend{"xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
		{{end}}
		`)
	p := Person{"lee", []string{"aaaa@qq.com", "bbbbb@qq.com"}, []*Friend{&f1, &f2}}
	t.Execute(rw, p)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

}
