package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// 读客户端发送的信息
	for {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("读到的buf=", buf)
	}
}

func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	// 一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接我...........")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept()n err=", err)
			return
		}
		// 一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
