package main

import (
	"flag"
	"fmt"
	"strconv"
)

var levin int

func main() {
	// 初始化变量的方式
	var hp int32 = 100;
	hp1 := 100;
	var hp2 = 100;
	var a, b = 1, 2
	//交换值
	a, b = b, a
	fmt.Println(hp, hp1, hp2)
	fmt.Println(a, b)
	//自动推导类型
	str := "I'am string"
	fmt.Println(str)
	// type 用法
	var o IntSlice = []int{1, 2, 3}
	fmt.Println("IntSlice len is :", o.Len())
	//多行字符串
	const con = `
		Hello World
	`
	//构造对象
	p := Person{
		name:    "xiaoliwen",
		age:     0,
		address: "",
	}
	var name, _ = p.getNameAddress()
	fmt.Println("Persion Filed for name value is :", name)
	//强制转换string成int
	var i, _ = strconv.Atoi("920712")
	fmt.Println("string convert int :", i)
	//获取内存地址
	ptri := &i
	//输出指针类型
	fmt.Printf("ptri type: %T \n", ptri)
	//指针取值
	fmt.Println(*ptri)
	// 使用指针交换两个变量
	var x, y = 1, 2
	swap(&x, &y)
	var mode = flag.String("sql", "", "process mode")
	flag.Parse()
	fmt.Println(*mode)
	var day = Weekday(Sunday)
	fmt.Println(day.String())
}

type Weekday int

const (
	Sunday = 2 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	}
	return "N/A"
}

func getWeek(week Weekday) {
	fmt.Println(week)
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) getNameAddress() (string, string) {
	return p.name, p.address
}

func (p Person) getAge() int {
	return p.age
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp

}
