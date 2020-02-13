package main

import (
	"container/list"
	"flag"
	"fmt"
	"strconv"
	"sync"
)

var levin int

/**
 * 关于Slice的常用操作
 * 切片扩容是Cap*2
 */
func sliceExample() {
	//固定长度初始化数组
	var s1 [5]string = [5]string{"A", "B", "C", "D", "E"}
	//动态长度初始化数组
	var s2 = [...]string{"A", "B", "C", "D", "E"}
	//遍历数组
	for i, v := range s1 {
		fmt.Printf("索引%d,的值是%s \n", i, v)
	}
	fmt.Printf("s1与s2是否相等? %t \n", s1 == s2)
	// 切片截取
	fmt.Printf("s1[1:2] is :%s \n", s1[1:2])
	// 切片中追加元素,第一个参数不可以是定长切片,第二参数如果是切片需要加...
	var s3 = append(s1[:5], s2[:5]...)
	fmt.Printf("append value is :%s \n", s3)
	copy(s1[0:3], s2[2:5])
	fmt.Printf("copy value is :%s \n", s1)
	//make函数构造切片
	var makeS = make([]int, 3, 10)
	makeS[1] = 1
	fmt.Printf("makeS is: %s \n", makeS)
}

func mapExample() {
	//初始化一个定长10的map,扩容的话每次加1
	var m1 map[string]string = make(map[string]string, 10)
	m1["name"] = "xiaoliwen"
	//初始化一个Map 并且填充数据
	var m2 = map[string]string{"name": "chenchen", "address": "mudanjiang"}
	//删除指定key
	delete(m2, "name")
	//判断值
	if _, ok := m1["address"]; ok == false {
		fmt.Println("没找到key")
	}
	for k, v := range m1 {
		fmt.Printf("m1 key is %s,value is %s \n", k, v)
	}
	for k, v := range m2 {
		fmt.Printf("m2 key is %s,value is %s \n", k, v)
	}
	//清空Map,不用担心效率
	m2 = make(map[string]string)
}
func syncMapExample() {
	//初始化一个同步的Map
	//var syncM = new(sync.Map)
	var syncM = sync.Map{}
	//向Map填充数据
	syncM.Store("name", "xiaoliwen")
	syncM.Store("address", "jixi")
	syncM.Store("age", "28")
	//便利Map并且删除key为address的
	syncM.Range(func(key, value interface{}) bool {
		if key != "address" {
			fmt.Printf("Sync.Map key is %s, value is %s \n", key, value)
		} else {
			syncM.Delete(key)
		}
		return true
	})
}

/**
 * List常用操作
 * list没有类型限制,这样既带来了便利,但是也有一些问题,如果向list中填充了interface{},那么在取出的时候转换为其他类型会宕机
 */
func listExample() {
	//初始化list
	var l1 = list.New()
	//添加元素到尾部
	var temp = l1.PushBack("A")
	//添加元素到头部
	l1.PushFront("B")
	//删除元素,是一个指针类型
	l1.Remove(temp)
	//遍历list
	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Printf("list value is :%s \n", i.Value)
	}
}
func main() {
	fmt.Println(nil==nil)
}
func main1() {
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
	fmt.Printf("IntSlice len is : %d, cap is : %d , Type is : %T", o.Len(), cap(o), o)
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
	/**
	 * sliceExample实例
	 */
	sliceExample()
	/**
	 * map实例
	 */
	mapExample()
	/**
	 * syncMap实例
	 */
	syncMapExample()
	/**
	 * list实例
	 */
	listExample()
	//nil值 nil 是 map、slice、pointer、channel、func、interface的零值
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
