package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jiangboLee/go_web_study/leechitchat/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var (
	Config Configuration
	logger *log.Logger
)

//快速打印构造函数
func P(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	/*
	   // 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	   // os.O_CREATE|os.O_APPEND
	   // 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	   // os.O_RDONLY // 只读
	   // os.O_WRONLY // 只写
	   // os.O_RDWR // 读写
	   // os.O_APPEND // 往文件中添建（Append）
	   // os.O_CREATE // 如果文件不存在则先创建
	   // os.O_TRUNC // 文件打开时裁剪文件
	   // os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	   // os.O_SYNC // 以同步I/O的方式打开


	   //第3个是打开时的属性
	*/
	file, err := os.OpenFile("leechitchat", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开log文件失败", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("不能打开文件", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("不能从文件读取json", err)
	}
}

//错误信息
func error_message(rw http.ResponseWriter, req *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(rw, req, strings.Join(url, ""), 302)
}

//得到这个cookie的session
func session(rw http.ResponseWriter, req *http.Request) (sess data.Session, err error) {
	cookie, err := req.Cookie("_cookie")
	P(cookie)
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		P("找到的session")
		P(sess)
		if ok, _ := sess.Check(); !ok {
			P("okokokokokoko?????????")
			err = errors.New("无效session")
		}
	}
	return
}

// parse html
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

//logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func Version() string {
	return "0.1"
}
