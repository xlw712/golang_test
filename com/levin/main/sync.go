package main

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var sLock *sync.Mutex
var rwLock *sync.RWMutex
var wg *sync.WaitGroup

var noneChan chan string

/**
 * 这个方法是用来测试单向通道的
 * 这个方法将一个双向通道,进行消费逻辑,然后隐私的将通道转换为单向只写通道,避免了其他地方的消费
 */
func chan2onlyRead() chan<- string {
	noneChan = make(chan string)
	go func() {
		for value := range noneChan {
			fmt.Println(value)
		}
	}()
	return noneChan
}
/**
 * 运用select 来处理IO多路复用,select中的case处理的是一个io操作,并且测试了一下超时处理
 */
func selectExample() {
	ch := make(chan int)
	//quit通道用来当做一个退出标识
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(time.Second * 3):
				fmt.Println("超时了")
				quit <- true
			}
		}

	}()
	for i:=0;i<=10 ;i++  {
		ch<-i
		time.Sleep(time.Second)
	}
	//当消费者超时的时候,才会向通道中写入一条信息,这个是时候才能解除阻塞
	<-quit
}
func main() {
	selectExample()
	onlyWrite := chan2onlyRead()

	onlyWrite <- "1111"
	onlyWrite <- "2222"
	onlyWrite <- "3333"

	var input string
	for input != "exit" {
		fmt.Scanln(&input)
	}

}
func main1() {
	// chan 如果没有缓冲,那么进出都会阻塞
	noneBufferChan := make(chan string)
	//BufferChan := make(chan struct{}, 10)
	//onlyReadChan := make(<-chan struct{})
	//onlyWriteChan := make(chan<- struct{})
	//Product
	sLock = new(sync.Mutex)
	wg = new(sync.WaitGroup)
	fmt.Printf("系统Cput个数:%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 5; i++ {
		/**
		 *这里要用temp,也就是闭包,给匿名函数增加入参,因为是值传递,所以每次循环的时候都会copy一个临时变量到匿名函数中去
		 */
		go func(temp int, c chan string) {
			c <- strconv.Itoa(temp)
		}(i, noneBufferChan)
	}
	//Consumer
	count := 0
	wg.Add(1)
	go func(c chan string, wg *sync.WaitGroup) {
		wg.Wait()
		//for遍历通道在没有元素时会阻塞
		for value := range c {
			fmt.Printf("消费者消费了,%s \n", value)
			count++
		}

	}(noneBufferChan, wg)
	/**
	 * 1.go run sync.go -name 111
	*  1.go run sync.go --help
	*/
	//这个返回参数是一个指针
	name := flag.String("name", "肖力文", "姓名")
	var nameVar string
	flag.StringVar(&nameVar, "nameVar", "变量肖力文", "变量姓名")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(nameVar)
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	var intput string
	fmt.Println("消费者等待,接受用户的输入:")
	fmt.Scanln(&intput)
	fmt.Printf("用户输入的是:%s\n", intput)
	wg.Done()
	fmt.Println("等待关闭指令")
	fmt.Scanln(&intput)
	if intput != "exit" {
		panic("关闭指令错误,应该是exit")
	}

}
