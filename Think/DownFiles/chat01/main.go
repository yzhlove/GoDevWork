package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	//var u1  = "http://ddns.10086.fund:23339/"
	//u2 := "http://ddns.10086.fund:23339/03-%E4%BB%8E0%E5%BC%80%E5%A7%8B%E5%AD%A6%E6%9E%B6%E6%9E%84/01-%E5%BC%80%E7%AF%87%E8%AF%8D%20%281%E8%AE%B2%29/"
	//u3 := "http://ddns.10086.fund:23339/03-%E4%BB%8E0%E5%BC%80%E5%A7%8B%E5%AD%A6%E6%9E%B6%E6%9E%84/01-%E5%BC%80%E7%AF%87%E8%AF%8D%20%281%E8%AE%B2%29/00%E4%B8%A8%E5%BC%80%E7%AF%87%E8%AF%8D%E4%B8%A8%E7%85%A7%E7%9D%80%E5%81%9A%EF%BC%8C%E4%BD%A0%E4%B9%9F%E8%83%BD%E6%88%90%E4%B8%BA%E6%9E%B6%E6%9E%84%E5%B8%88%EF%BC%81.html"

	u4 := "http://ddns.10086.fund:23339/03-%E4%BB%8E0%E5%BC%80%E5%A7%8B%E5%AD%A6%E6%9E%B6%E6%9E%84/01-%E5%BC%80%E7%AF%87%E8%AF%8D%20%281%E8%AE%B2%29/00%E4%B8%A8%E5%BC%80%E7%AF%87%E8%AF%8D%E4%B8%A8%E7%85%A7%E7%9D%80%E5%81%9A%EF%BC%8C%E4%BD%A0%E4%B9%9F%E8%83%BD%E6%88%90%E4%B8%BA%E6%9E%B6%E6%9E%84%E5%B8%88%EF%BC%81.html"

	r, err := http.Get(u4)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	f, err := os.Create("web.html")
	if err != nil {
		panic(err)
	}
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
	f.Close()
}
