package main

import (
	"errors"
	"fmt"
	"sort"
)

type Cat interface {
	say()
	run()
	catchFish()
}
type BlackCat struct {
}

func (b BlackCat) say() {
	fmt.Println("喵")
}

func (b BlackCat) run() {
	fmt.Println("奔跑")
}

func (b BlackCat) catchFish() {
	fmt.Println("捕鱼")
}

type Student struct {
	name string
	age  int
}
type Students []*Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].age > s[j].age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type CustomerError struct {
}

func (*CustomerError) Error() string {
	return "这是一个自定义异常"
}

func main() {
	var cat Cat
	cat = new(BlackCat)
	cat.say()
	cat.run()
	cat.catchFish()
	//断言
	var x interface{}
	x = "xiaoliwen"
	if value, ok := x.(string); ok == true {
		fmt.Println("类型匹配:" + value)
	}
	//依靠排序结构排序,需要构造一个Student切片
	s := Students{&Student{"小明", 1}, &Student{"小红", 22}, &Student{"小刚", 19}, &Student{"小李", 5}}
	sort.Sort(s)
	for _, v := range s {
		fmt.Printf("%+v\n", v)
	}
	//切片排序
	sliceStudent := []*Student{
		{"老大", 12},
		{"老二", 9},
		{"老三", 3},
	}
	sort.Slice(sliceStudent, func(i, j int) bool {
		return sliceStudent[i].age < sliceStudent[j].age
	})
	for _, v := range sliceStudent {
		fmt.Printf("%+v\n", v)
	}
	// 异常接口实例
	var error error
	error = new(CustomerError)
	fmt.Printf("这个异常的信息是:"+error.Error())

	errors.New("new Error mode")
}
