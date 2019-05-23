package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "39.106.33.201:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tcp has connect")
	defer conn.Close()
	go func() {
		fmt.Fprint(conn, "POST /cgi-bin/hello HTTP/1.1\r\n",
			"Host:xyzh.daymenu.cn\r\n",
			"Content-type:application/x-www-form-urlencoded\r\n",
			"\r\n",
			"name=hanjian&sex=nan") 
	}()
	fmt.Println("request has send")
	html, err := ioutil.ReadAll(conn)
	fmt.Println(string(html))

}
