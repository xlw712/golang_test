package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name            string   `json:"type" id:"712"`
	age             int      `年龄`
	history_address []string `历史地址`
}
type B func()

type Color int

const (
	Red         Color        = 1 << iota // 1
	Black                    = 100       //100
	White                                //100
	Blue, Yello = iota, iota             //3,3
	Green       = iota                   //4
)

func printReflect() {
	var vars []interface{}
	var a int
	var b string
	var c bool
	var d A
	var e B
	var f *A
	var g [3]int
	var h [] int
	var i Color
	vars = []interface{}{a, b, c, d, e, f, g, h, i}
	for index, t := range vars {
		//返回一个Type类型
		typeOf := reflect.TypeOf(t)
		//打印类型的名字和种类
		fmt.Printf("第 %d 个 变量的类型是 %s ,种类是 %s \n", index+1, typeOf.Name(), typeOf.Kind())
	}
}
func main() {
	printReflect()
	fmt.Println("--------1------")
	var i int
	typeOf := reflect.TypeOf(&i)
	fmt.Println(typeOf.Name(), typeOf.Kind())
	//reflect.Elem方法在获取一个指针类型时,会隐私的对这个指针类型做一个*操作
	typeElem := typeOf.Elem()
	fmt.Println(typeElem.Name(), typeElem.Kind())
	fmt.Println("------2--------")
	var a A = A{"xiaoliwen", 28, []string{"hrb", "beijing", "jixi"}}
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		// 打印字段名称,字段Tag,和字段Tag中的值
		// Tag的格式为`key1:"value1" key2:"value2"`
		fmt.Println(filed.Name, filed.Tag, filed.Tag.Get("json"), filed)
	}
	fmt.Println("------3--------")
	/**
	 * 如果想通过反射修改一个对象,那么需要获取指针,并且使用Elem函数进行指针的解引用
	 * http://c.biancheng.net/view/5131.html
	 */
	v := reflect.ValueOf(&a).Elem()
	typeOfT := v.Type()
	for i := 0; i < typeOfT.NumField(); i++ {
		f := v.Field(i)
		//注意结构体中的字段只有是public才能反射
		if f.CanSet() {
			//通过反射修改了字段的值
			f.SetString("Levin")
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
	}
}
