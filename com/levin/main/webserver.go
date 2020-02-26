package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/welcome",welcome)
	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func welcome(w http.ResponseWriter, r *http.Request){
	 if content,error:=ioutil.ReadFile("./welcome.html");error==nil{
		 w.Write(content)
	 } else {
		fmt.Println("出错了")
	 }

}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "C语言中文网")
}
