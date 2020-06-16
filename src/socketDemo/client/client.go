package main

import "net"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 09:53
 *@version v1.0
 */
func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:9000")
	conn, _ := net.DialTCP("tcp", nil, addr)
	conn.Write([]byte("客户端发送信息"))
	conn.Close()

}
