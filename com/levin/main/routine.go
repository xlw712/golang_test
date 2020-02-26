package main

import (
	"fmt"
	"time"
)

func routine(str string) string {
	for {
		time.Sleep(time.Second)
		fmt.Println(str)
	}

	return str
}

func main() {
	go routine("Hello Routine")
	var input string
	//main方法也是启动的一个goroutine,如果main方法结束,那么其他goroutine也会结束,并且目前没有什么好方法获得routineID
	fmt.Scanln(&input)

}
