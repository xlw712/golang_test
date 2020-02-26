package main

/**
 * 文件的执行顺序 全局变量--->init函数--->main函数
 */
import (
	"fmt"
	m "golang_test/com/levin/mode"
)

var s1 int64 = s()

func init() {
	fmt.Println("init")
}
func s() int64 {
	fmt.Println("function s() --->")
	return 1
}

type Class interface {
	Do()
}

var (
	factoryByName = make(map[string]func() Class)
)

func Register(name string, factory func() Class) {
	factoryByName[name] = factory
}

type ClassImpl struct {
}

func (c ClassImpl) Do() {
	fmt.Println("工厂方法实现!!")
}
func Create(name string) Class {
	if value, ok := factoryByName[name]; ok {
		return value()
	} else {
		panic("name not found")
	}
}

// 单例模式
type Tool struct {
	value int
}

func main() {
	person := new(m.Person)
	fmt.Println(person)
	person.SayHello()
	f := func() Class {
		return new(ClassImpl)
	}
	Register("First Factory Impl", f)
	Create("First Factory Impl").Do()

}
