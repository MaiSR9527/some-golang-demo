package main

import (
	"fmt"
	"net"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/16 08:53
 *@version v1.0
 */
func main() {
	//1.创建服务器地址
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:9000")
	//2.创建监听器
	listener, _ := net.ListenTCP("tcp4", addr)
	//3.通过监听器获取客户端传递过来的数据
	//阻塞式
	conn, err := listener.Accept()
	if err != nil {
		panic(err.Error())
	}
	//4.转换数据
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	fmt.Println("服务器端接收的数据：", string(b[:n]))
	//5.关闭连接
	err = conn.Close()
	if err != nil {
		panic(err.Error())
	}
}
