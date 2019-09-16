package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dail err=", err)
		return
	}
	defer conn.Close()
	fmt.Println("coon success", conn)
	_, err = conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("set err= ", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err= ", err)
		return
	}
	fmt.Println("name= ", r)
	fmt.Println("操作成功")
}
