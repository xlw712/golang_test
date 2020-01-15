package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 主要用来熟悉golang函数定义
 */
func main() {
	str1 := "a"
	str2 := "b"
	isExistsArray("a", []*string{&str1, &str2}...)
	// 函数标识参数化
	test(1, 2, mul)
	test(1, 2, add)
	//函数变量化
	var f func(int)
	f = fire
	f(1)
	//链式处理
	strList := []string{"  hello", "List"}
	//匿名函数
	chain := []func(str string) string{
		strings.ToUpper,
		strings.TrimSpace,
		concat,
	}
	StringProcess(strList, chain)
	//也可以像下面这么写
	//StringProcess(strList, []func(str string) string{
	//	strings.ToUpper,
	//	strings.TrimSpace,
	//	concat,
	//})

	fmt.Println(strList)
	//匿名函数调用
	func(str string) {
		fmt.Println(str)
	}("我是一个匿名函数")
	//接口用法
	audi := Audi{}
	bmw := Bmw{}
	var car Car
	car = audi
	car.run(100)
	car = bmw
	car.run(120)
	checkInterface(audi)
}

/**
 * 一个普通的函数,判断切片中是否存在对应的字符串
 */

func isExistsArray(str string, arr ...*string) (b bool, log string) {
	for _, value := range arr {
		if *value == str {
			return true, "存在字符串"
		}
	}

	return false, "不存在字符串"
}

/**
 * 测试函数类型参数化
 */
type CalculateType func(a, b int) (result int)

func add(a, b int) (result int) {
	fmt.Println(strconv.Itoa(a), "+", strconv.Itoa(b))
	result = a + b
	return
}
func mul(a, b int) (result int) {
	fmt.Println(strconv.Itoa(a), "*", strconv.Itoa(b))
	result = a * b
	return
}

func test(a, b int, calculate CalculateType) {
	fmt.Println(calculate(a, b))
}
func (c CalculateType) toString() {
	fmt.Printf("%T \n", c)
}

/**
 * 函数变量化
 */
func fire(i int) {
	fmt.Println("fire")
}

/**
 * 字符串链式处理
 */
func StringProcess(strlist []string, chain []func(string) string) {
	for index, str := range strlist {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		strlist[index] = result
	}
}
func concat(str string) (string) {
	return str + "!!!"
}

/**
 * 函数类型实现接口
 */
type Car interface {
	run(km int)
}
type Audi struct{}
type Bmw struct{}

func (Audi) run(km int) {
	fmt.Println("奥迪能跑 " + strconv.Itoa(km) + "km")
}
func (Bmw) run(km int) {
	fmt.Println("宝马能跑 " + strconv.Itoa(km) + "km")
}
func checkInterface(t interface{},i interface{}) {
	var car Car = Audi{}
	fmt.Printf("%T", car)
	if _, ok := t.(Car); ok {
		fmt.Println("实现了Car接口")
	} else {
		fmt.Println("没有实现Car接口")
	}
}
