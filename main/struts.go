package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type Father struct {
	name string
	age  int
}

func (f Father) say() {
	fmt.Println("我是父亲,我的名字叫" + f.name)
}

type Son struct {
	Father
}

func (s Son) say() {
	fmt.Println("我是儿子,我父亲的名字叫", s.Father.name)
}

// Json结构
// 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int // 容量
}

func getJson() []byte {
	// 完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		// 电池参数
		Battery: Battery{
			2910,
		},
		// 是否有指纹识别
		HasTouchID: true,
	}
	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw)

	return jsonData

}
func main() {
	/**
	 * 继承例子 Golang中不建议用继承,用组合更好
	 */
	var f Father = Father{name: "刘备", age: 50}
	f.say()
	var s Son = Son{f}
	s.say()
	/**
	 * 使用匿名Struts解析Json结构
	 */
	jsonByte := getJson()
	fmt.Println("构造Jsons数据,格式为" + string(jsonByte))
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	if error := json.Unmarshal(jsonByte, &screenAndTouch); error == nil {
		fmt.Printf("反序列化之后的格式是,%+v\n", screenAndTouch)
	} else {
		fmt.Println("error:" + error.Error())
	}
	runtime.SetFinalizer(&screenAndTouch, func(s *struct {
		Screen
		HasTouchID bool
	}) {
		jsonData, _ := json.Marshal(s)
		fmt.Printf("Struts对象被回收,格式为Json:%s\n", string(jsonData))
	})
	for i := 0; i <= 10; i++ {
		runtime.GC()
	}

}
