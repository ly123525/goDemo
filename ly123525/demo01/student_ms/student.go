package main

import "fmt"

// Student 结构体
type Student struct {
	name  string
	age   int
	id    int64
	class string
}

// NewStudent 是一个创建新学生对象构造函数
func NewStudent(name string, age int, id int64, class string) *Student {
	return &Student{
		name:  name,
		age:   age,
		id:    id,
		class: class,
	}
}

// StudentMgr 学生信息管理的结构体
type StudentMgr struct {
	AllStudents []*Student
}

// NewStudentMgr 新的学生信息结构体对象
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		AllStudents: make([]*Student, 0, 100),
	}
}

// AddStudent 添加学生
func (s *StudentMgr) AddStudent(stu *Student) {
	for _, v := range s.AllStudents {
		if v.name == stu.name {
			fmt.Printf("姓名为%s的学生已经存在\n", stu.name)
			return
		}
	}
	s.AllStudents = append(s.AllStudents, stu)
}

// EditStudent 修改学生
func (s *StudentMgr) EditStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			s.AllStudents[index] = stu
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在", stu.name)
}

// ShowStudent 展示学生
func (s *StudentMgr) ShowStudent() {
	for _, v := range s.AllStudents {
		fmt.Printf("name: %s, age: %d, id: %d, class: %s\n", v.name, v.age, v.id, v.class)
	}
}

// DeleteStudent 删除学生
func (s *StudentMgr) DeleteStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if v.name == stu.name {
			s.AllStudents = append(s.AllStudents[:index], s.AllStudents[index+1:]...)
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在", stu.name)
}
