package main

import "fmt"
import "encoding/json"

type person struct {
	Name string
	Age  int
}

type person2 struct {
	Name string `json:"name" db:"name" ini:"name"` //不同应用场景 该变量取不同的名字
	Age  int    `json:"age"`
}

func newPerson(name string, age int) person {
	return person{Name: name, Age: age}
}

func (p person) angry() {
	fmt.Println(p.Name, "生气了")
}

func main() {
	var p1 = newPerson("小红", 20)
	p1.angry()
	p2 := person{"小明", 9000}
	//序列化
	j, err := json.Marshal(p2)
	if err != nil {
		fmt.Printf("marshal faild, err:%v", err)
	}
	fmt.Println(j, err, string(j))
	//person 属性小写时： [123 125] <nil> {}
	//person 属性大写时： [123 34 78 97 109 101 34 58 34 229 176 143 230 152 142 34 44 34 65 103 101 34 58 57 48 48 48 125] <nil> {"Name":"小明","Age":9000}

	p3 := person2{Name: "li", Age: 30}
	j2, err2 := json.Marshal(p3)
	println(j2, string(j2), err2)

	//反序列化
	var p4 person
	str := `{"name":"li","age":130}`
	json.Unmarshal([]byte(str), &p4) //传入指针 在Unmarshal中修改p3的值
	fmt.Println(p4)
	fmt.Printf("%v", p4)
}
