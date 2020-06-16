package main

import (
	"fmt"
	"net"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 10:05
 *@version v1.0
 */
func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:9000")
	conn, _ := net.DialTCP("tcp", nil, addr)
	conn.Write([]byte("客户端发送信息1"))
	b := make([]byte, 1024)
	count, _ := conn.Read(b)
	fmt.Println("客户端接收信息：", string(b[:count]))
	conn.Close()
}
