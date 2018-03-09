package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIp   string `json:"serverIp"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	//json解析
	/*var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)*/
	//生成json
	var s Serverslice
	s.Servers = append(s.Servers, Server{"shanghai", "127.0.0.1"})
	s.Servers = append(s.Servers, Server{"zhejiang", "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json error: ", err)
	}
	fmt.Println(string(b))
}
