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
	addr, _ := net.ResolveTCPAddr("tcp4", ":9999")
	lis, _ := net.ListenTCP("tcp4", addr)

	for {
		conn, _ := lis.Accept()
		go func() {
			for {
				b := make([]byte, 512)
				//读取数据
				count, _ := conn.Read(b)

				arrStr := strings.Split(string(b[:count]), "-")
				user.Username = arrStr[0]
				user.OtherUsername = arrStr[1]
				user.Msg = arrStr[2]
				user.ServerMsg = arrStr[3]
				userMap[user.Username] = conn
				if v, ok := userMap[user.OtherUsername]; ok && v != nil {
					user.ServerMsg = ""
					n, e := v.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
					if n == 0 || e != nil {
						conn.Close()
						delete(userMap, user.OtherUsername)
						break
					}
				} else {
					user.ServerMsg = "对方不在线"
					n, e := conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
				}
			}
		}()
	}
}
