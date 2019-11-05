package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HttpStart(port int) {
	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("监听失败：", err.Error())
	}
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("打印Header参数列表：")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	fmt.Println("打印Form参数列表：")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
	//验证用户名密码，如果成功则header里返回session，失败则返回StatusUnauthorized状态码
	w.WriteHeader(http.StatusOK)
}

func main() {

	HttpStart(1234)

}
