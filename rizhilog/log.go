package main

import (
	"flag"
	"strings"
	"strconv"
	"fmt"
	"net/url"
	"math/rand"
	"time"
	"os"
)

type resource struct {
	url string
	target string
	start int
	end int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		 url: "http://www.ljbniubi.top",
		 target: "",
		 start: 0,
		 end: 0,
	}
	r2 := resource{
		"http://www.ljbniubi.top/{$id}.html",
		"{$id}",
		1,
		12,
	}
	r3 := resource{
		"http://www.ljbniubi.top/movie/{$id}.html",
		"{$id}",
		1,
		12222,
	}
	res = append(append(append(res, r1), r2), r3)
	return res
}
func buildurl( res []resource) []string {
	var list []string
	for _, value := range res {
		if len(value.target) == 0 {
			list = append(list, value.url)
		} else {
			for i:=value.start; i<value.end; i++  {
				urlstr := strings.Replace(value.url, value.target, strconv.Itoa(i),-1)
				list = append(list, urlstr)
			}
		}
	}
	return list
}

func makelog(current, refer, ua string) string {
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramsStr := u.Encode()

	logTemplate := "127.0.0.1 http://www.ljbniubi.top?" + paramsStr
	return logTemplate
}
func randInt(start, end int) int {
	r := rand.New( rand.NewSource(time.Now().UnixNano()))
	if end < start {
		 return end
	}
	return r.Intn(end - start) + start
}
func main() {
	total := flag.Int("total", 100, "how many rows")
	filePath := flag.String("filepath","/Users/ljb48229/Desktop/log", "file path")
	flag.Parse()
	fmt.Println(*total, *filePath)


	res := ruleResource()
	list := buildurl(res)
	//fmt.Println(list)
	logStr := ""
	for i:=0; i<=*total ; i++ {

		currentUrl := list[ randInt(0, len(list) - 1) ]
		refer := list[ randInt(0, len(list) - 1)]
		logStr += makelog(currentUrl, refer, "hahaha" ) + "\n"
	}
	file, err := os.OpenFile(*filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {

	}
	file.Write([]byte(logStr))
	file.Close()
}
