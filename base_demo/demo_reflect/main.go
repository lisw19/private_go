package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type cat struct {
	color string
}

// TypeOf
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
	// 打印类型和所属种类
	fmt.Printf("type_name: %v type_kind: %v \n", v.Name(), v.Kind())
}

// ValueOf
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d \n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is Float32, value is %s \n", v.Float())
	case reflect.Float64:
		fmt.Printf("type is Float64, value is %s \n", x)
	case reflect.Struct:
		fmt.Printf("type is Struct, value is %s \n", v)
	}
}

// 通过反射设置变量的值
func relectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
}

type RedisConfig struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
	Db   int    `ini:"db"`
}

type MysqlConfig struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
	Db   int    `ini:"db"`
}

type Config struct {
	RedisConfig `ini:"redis"`
	MysqlConfig `ini:"mysql"`
}

func loadIni(iniFile string, v interface{}) (err error) {
	//0.参数校验， v参数必须是指针类型，因为需要修改其值，且其必须是以结构体
	t := reflect.TypeOf(v)
	fmt.Println(t.Name(), '|', t.Kind(), '|', t.Elem())
	if t.Kind() != reflect.Ptr {
		//err = fmt.Errorf("param error, v type error, must be ptr") //格式化输出 发挥error类型
		// 创建一个错误
		err = errors.New("param error, v must by ptr")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("param error, v must by struct ptr")
		return err
	}
	//1.加载ini文件
	bd, err := ioutil.ReadFile(iniFile)
	if err != nil {
		return
	}
	var struct_name string
	//2.按行读取文件
	// 把文件内容转为字符串
	lineSlice := strings.Split(string(bd), "\n")
	for idx, line := range lineSlice {
		//2.1如果是注释就跳过
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2 如果是括号，表示是一个节点
		if strings.HasPrefix(line, "[") {
			if line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			section_name := strings.TrimSpace(line[1 : len(line)-1])
			if len(section_name) < 1 {
				err = fmt.Errorf("line:%d, title is null", idx+1)
				return
			}
			// 根据结构体节点名反射值
			for i := 0; i < t.Elem().NumField(); i++ {
				field_name := t.Elem().Field(i).Tag.Get("ini")
				if section_name == field_name {
					// 找到对应机构体名
					struct_name = field_name
					fmt.Printf("找到ini文件中节点%s对应嵌套结构体%s\n", section_name, struct_name)
				}
			}
		} else {
			//空行跳过
			if len(line) == 0 {
				continue
			}
			// 以'='号拆分出值
			if strings.Index(line, "=") < 0 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			// 根据structName 去 结构里把结构体字段取出来
			value := reflect.ValueOf(v)
			structObj := value.Elem().FieldByName(struct_name)
			fmt.Println(structObj.Kind())
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("传入的v中， %s需是结构体", structObj)
				return
			}

		}
	}
	//2.3 其他就是键值对
	return
}

func main() {
	var a = 2.12
	reflectType(a)
	reflectValue(a)

	var b int64
	b = 100
	reflectType(b)
	reflectValue(b)

	var c cat
	c = cat{"红色"}
	reflectType(c)
	reflectValue(c)

	var d int64 = 3
	relectSetValue(&d)
	fmt.Println(d)

	var config Config
	err := loadIni("/Users/lishiwei/Workspace/GO/src/"+
		"private_go/base_demo/demo_reflect/config.ini", &config)
	if err != nil {
		fmt.Printf("load ini error: %v \n", err)
	}

}
