package main

import (
	"fmt"
	"net"
	"strings"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 10:23
 *@version v1.0
 */
type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	userMap = make(map[string]net.Conn)
	user    = new(User)
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:9000")
	listener, _ := net.ListenTCP("tcp4", addr)
	for {
		accept, _ := listener.Accept()
		go func() {
			b := make([]byte, 1024)
			read, _ := accept.Read(b)
			array := strings.Split(string(b[:read]), "-")
			if user == nil {
				user = new(User)
			}
			user.Username = array[0]
			user.OtherUsername = array[1]
			user.Msg = array[2]
			user.ServerMsg = array[3]

			if v, ok := userMap[user.OtherUsername]; ok && v != nil {
				s := fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)
				n, err := v.Write([]byte(user.Msg))
				if n <= 0 || err != nil {
					delete(userMap, user.OtherUsername)
					accept.Close()
					v.Close()
					return
				}
			} else {
				user.ServerMsg = "对方不在线"
				accept.Write()
			}
		}()
	}
}
