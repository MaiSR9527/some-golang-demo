package main

import (
	"fmt"
	"net"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 09:59
 *@version v1.0
 */
func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:9000")
	listener, _ := net.ListenTCP("tcp4", addr)
	//阻塞式
	conn, err := listener.Accept()
	if err != nil {
		panic(err.Error())
	}
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	fmt.Println("服务器端接收的数据：", string(b[:n]))
	conn.Write(append([]byte("server:"), b[:n]...))
	//关闭连接
	err = conn.Close()
	if err != nil {
		panic(err.Error())
	}

}
