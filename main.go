package main

import (
	"container/list"
	"flag"
	"fmt"
	"strconv"
	"sync"
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
	// 接收启动参数 go run --sql "select * from t"
	var mode = flag.String("sql", "", "process mode")
	flag.Parse()
	fmt.Println(*mode)
	// 定义枚举类型
	var day = Weekday(Sunday)
	fmt.Printf("%T \n", day)
	// 构造数组 遍历数组
	var arrInt = []int{1, 2, 3, 4}
	// 使用切片获取数据 数据3,4 下标从0开始 前包含后不包含
	fmt.Println(arrInt[2:])
	// 向数组中追加数据
	arrInt = append(arrInt, 6, 7, 8)
	// 合并两个切片注意一定要加...
	arrInt = append(arrInt[0:2], arrInt[3:7]...)
	for k, v := range arrInt {
		fmt.Println(k, v, cap(arrInt))
	}
	//初始化map
	var map1 map[string]string
	//如果不执行make会报错,因为map1是一个Nil
	map1 = make(map[string]string, 10)
	map1["name"] = "xiaoliwen"
	map1["country"] = "china"
	if _, ok := map1["noExists"]; ok == false {
		fmt.Println("没找到key")
	}
	delete(map1, "name")
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	//初始化syncmap
	conMap := new(sync.Map)
	conMap.Store("who is Map type?", "I'am concurrent Map")
	conMap.Store("I'am crazy", "yes")
	fmt.Println(conMap.Load("who is Map type?"))
	conMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, "->", value)
		return true
	})
	//初始化list
	var list = list.New()
	list.PushBack("A")
	list.PushFront("B")
	list.PushBack("C")
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	//nil值
	var a1 string
	var a2 string
	fmt.Println(a1 == a2)
	// if
	if _, error := strconv.Atoi("9207d12"); error != nil {
		fmt.Println(error)
	}
	//遍历通道接收数据
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	//switch
	str1 := "hello"
	switch {
	case str1 == "hello":
		fmt.Println(str1)
		fallthrough
	case str1 != "xiaoliwen":
		fmt.Println("xiaoliwen")
	}
	// goto
	tmp_count := 10
	for {
		tmp_count--
		if tmp_count == 0 {
			goto exit
		}
	}
exit:
	{
		fmt.Println("退出循环")
	}

	var chinese = "我是中国人， I am Chinese"
	fmt.Println([]rune(chinese))
}

type Weekday int

const (
	Sunday Weekday = 2 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func getWeek(week Weekday) {
	fmt.Println(week)
}
func (w Weekday) String() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	}
	return "N/A"
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p IntSlice) add(v int)          { p = append(p, v) }

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
