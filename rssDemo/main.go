package main

import (
	_ "./matchers"
	"./search"
	"log"
	"os"
)

func init() {
	//将日志输出标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("why")
}
