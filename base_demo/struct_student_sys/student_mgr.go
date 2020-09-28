package main

import "fmt"

//结构体版的学生管理系统, 实现增删改查
type student struct {
	id   int64
	name string
}

// 学生管理对象
type studentM struct {
	allStudent map[int64]student
}

// 增
func (s studentM) addS() {
	// 根据用户输入内容 常见学生
	var (
		sid   int64
		sname string
	)
	fmt.Print("add请输入学号：")
	fmt.Scanln(&sid)
	fmt.Print("add请输入name：")
	fmt.Scanln(&sname)
	newStuden := student{
		id:   sid,
		name: sname,
	}
	//把学生加入管理对象中
	s.allStudent[sid] = newStuden
}

//删
func (s studentM) delS() {

}

//改
func (s studentM) changeS() {
	// 根据用户输入内容 常见学生
	var (
		sid      int64
		s_n_name string
	)
	fmt.Print("change请输入学号：")
	fmt.Scanln(&sid)
	d_student, studeng_statu := s.allStudent[sid]
	if studeng_statu {
		fmt.Println("找到学生", d_student)
		fmt.Print("change请输入修改后name：")
		fmt.Scanln(&s_n_name)
		d_student.name = s_n_name
		s.allStudent[sid] = d_student
		fmt.Println("修改成功", d_student)
		return
	} else {
		fmt.Print("change学生不存在")
		return
	}

}

//查
func (s studentM) showS() {
	for _, st := range s.allStudent {
		fmt.Println("序号", st.id, "   name", st.name)
	}
}
