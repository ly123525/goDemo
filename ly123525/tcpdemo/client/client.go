package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Printf("client dial suc=%v\n", conn)
	// 客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入

	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("rederString err=", err)
		}
		// 如果用户输入的是 exit就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
		}
		_, err = conn.Write([]byte(line + "\n"))

		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
}
