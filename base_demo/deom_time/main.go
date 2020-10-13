package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//时间戳转时间
	ret := time.Unix(1602554282, 0)
	fmt.Println(ret, ret.Hour())

	//时间差值
	fmt.Println(now.Add(time.Hour * 24))

	//定时 -- 一秒一次
	//timeer := time.Tick(time.Second)
	//for t := range timeer{
	//	fmt.Println(t)
	//}

	//时间格式化 Y(2006)-m(1)-d(2) H(3)：M(4)：S(5)
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02 03:04"))
	fmt.Println(now.Format("2006-01-02 03:04 PM")) //2020-10-13 10:17 AM

	//精确到秒
	fmt.Println(now.Format("2006/01/02 03:05:05.00")) //2020/10/13 10:02:02.41
	// str时间转time
	timeObj, err := time.Parse("2006-01-02", "2020-05-20")
	if err != nil {
		fmt.Println("time.Parse error")
		return
	}
	fmt.Println(timeObj.Date())

	//sub 时间间隔

	loc, errl := time.LoadLocation("Asia/Shanghai")
	if errl != nil {
		fmt.Println("LoadLocation error")
		return
	}
	//按照指定时区解析时间格式
	//nextyear, err2 := time.Parse("2006-01-02 15:04:05", "2020-10-13 20:00:00")
	nextyear, err2 := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-13 20:00:00", loc)
	if err2 != nil {
		fmt.Println("time.Parse error")
		return
	}
	fmt.Println(nextyear)
	now = now.UTC()
	jiange := nextyear.Sub(now)
	fmt.Println(jiange)
	fmt.Println("********1")
	time.Sleep(time.Second * 10)
	fmt.Println("********2")

}
