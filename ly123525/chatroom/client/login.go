package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ly123525/chatroom/common/message"
)

// 登录
func login(userID int, userPwd string) (err error) {
	//下一个就要开始定协议
	// fmt.Printf("userID = %d userPwd=%s\n", userID, userPwd)
	// return nil
	// 1. 链接到服务器
	conn, err := net.Dial("net", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	// 延迟关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd
	//4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. 把data赋值给mes.Data字段
	mes.Data = string(data)
	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 7. 到这个时候，data就是我们要发送的信息
	// 先把data的长度发送给服务器
	// 先获取到 data的长度-> 转换成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail", err)
		return
	}
	fmt.Println("客户端，发送消息的长度ok")
	return nil
}
