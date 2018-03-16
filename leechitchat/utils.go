package main

import (
	"./data"
	"encoding/json"
	"errors"
	"fmt"
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
	config Configuration
	logger *log.Logger
)

//快速打印构造函数
func p(a ...interface{}) {
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
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("不能从文件读取json", err)
	}
}

func version() string {
	return "0.1"
}
