package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//new/http server

//
func f1(res http.ResponseWriter, req *http.Request) {
	fileObj, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		// fmt.Println("read file failed ,err:", err)
		res.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	res.Write(fileObj) //强制转为字节类型的切片
}

func f2(res http.ResponseWriter, req *http.Request) {
	//对于GET请求，参数都放在了URL上(query param)，请求体中是没有数据的
	fmt.Println(req.URL)         ///xxx/?name=ads&age=18
	fmt.Println(req.URL.Query()) //自动识别URL中的query param  map[age:[18] name:[ads]]
	queryParam := req.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)                //ads 18
	fmt.Println(req.Method)               //GET
	fmt.Println(ioutil.ReadAll(req.Body)) //[] <nil>
	res.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
