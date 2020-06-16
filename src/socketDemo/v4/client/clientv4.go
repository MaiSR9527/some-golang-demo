package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 14:18
 *@version v1.0
 */
type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	user = new(User)
	wg   sync.WaitGroup
)

func main() {
	wg.Add(1)
	fmt.Println("请登录,输入用户名：")
	fmt.Scanln(&user.Username)
	fmt.Println("请输入要给谁发送消息")
	fmt.Scanln(&user.OtherUsername)
	addr, _ := net.ResolveTCPAddr("tcp4", ":9999")
	conn, _ := net.DialTCP("tcp4", nil, addr)
	go func() {
		fmt.Print("请输入:(只提示一次,以后直接输入即可)")
		for {
			fmt.Scanln(&user.Msg)
			if user.Msg == "exit" {
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
		}
	}()
	go func() {
		for {
			rb := make([]byte, 512)
			c, _ := conn.Read(rb)
			user2 := new(User)
			arrStr := strings.Split(string(rb[:c]), "-")
			user2.Username = arrStr[0]
			user2.OtherUsername = arrStr[1]
			user2.Msg = arrStr[2]
			user2.ServerMsg = arrStr[3]
			if user2.ServerMsg != "" {
				fmt.Println("\t\t\t服务器消息:", user2.ServerMsg)
			} else {
				fmt.Println("\t\t\t", user2.Username, ":", user2.Msg)
			}
		}
	}()
	wg.Wait()
}
