package main

import (
"fmt"
"log"
"net/http"
"strings"
)

func sahelName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello go!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sahelName)          //设置访问的路由
	err := http.ListenAndServe(":9000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("server run at http://127.0.0.1:9090")
}
