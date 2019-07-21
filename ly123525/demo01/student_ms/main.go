package main

import (
	"fmt"
	"os"
)

// 学生信息管理
// 添加学生
// 修改学生
// 删除学生
// 展示学生

func showMenu() {
	fmt.Println("学生信息管理系统！")
	fmt.Println("1. 添加学生")
	fmt.Println("2. 修改学生")
	fmt.Println("3. 展示学生")
	fmt.Println("4. 退出")
}

func userInput() *Student {
	var (
		name  string
		age   int
		id    int64
		class string
	)
	fmt.Println("请根据提示输入相关信息")
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入班级：")
	fmt.Scanln(&class)
	newStu := NewStudent(name, age, id, class)
	return newStu
}

func main() {
	stuMgr := NewStudentMgr()
	for {
		showMenu()
		var i int
		fmt.Print("请输入指令：")
		fmt.Scanln(&i)
		fmt.Println("输入的选项是：", i)

		switch i {
		case 1:
			newStu := userInput()
			stuMgr.AddStudent(newStu)
		case 2:
			newStu := userInput()
			stuMgr.EditStudent(newStu)
		case 3:
			newStu := userInput()
			stuMgr.DeleteStudent(newStu)
		case 4:
			stuMgr.ShowStudent()
		case 5:
			os.Exit(0)
		}
	}
}
