package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 这里我们循环的接收客户端发送的数据
	defer conn.Close()
	for {
		// 创建一个切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		// 2.如果客户端没有write[发送]， 那么协程就堵塞在这里
		fmt.Println("服务器在等待客户端发送信息")
		n, err := conn.Read(buf) // 从 conn读取
		if err != nil {
			fmt.Println("服务器的Read err=", err)
			return
		}
		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	//循环等待客户端来链接
	for {
		// 等待客户端来链接
		fmt.Println("等待客户端来链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err", err)
		} else {
			fmt.Printf("Accept() suc con =%v\n 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		// 这里准备一个协程, 为客户端服务
		go process(conn)
	}
}
