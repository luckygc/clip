package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/atotto/clipboard"
)

func main() {
	port := ":8081"
	fmt.Println("监听端口", port)
	printIp()
	http.HandleFunc("/push", push)
	http.HandleFunc("/pull", pull)        //设置访问的路由
	err := http.ListenAndServe(port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// func sayhello(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       //解析参数，默认是不会解析的
// 	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
// }

func pull(w http.ResponseWriter, r *http.Request) {
	text, _ := clipboard.ReadAll()
	fmt.Println("从pc获取剪贴板:", text)
	fmt.Fprintln(w, text) //这个写入到w的是输出到客户端的
}

func push(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	clipboard.WriteAll(r.PostFormValue("text"))
	fmt.Println("向pc推送剪贴板:", r.PostFormValue("text"))
	fmt.Fprintln(w, "success") //这个写入到w的是输出到客户端的
}

// func test(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //解析参数，默认是不会解析的
// 	fmt.Println(r.URL.Path)
// 	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
// }

func printIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
