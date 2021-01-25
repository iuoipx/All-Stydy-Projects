package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//net/http Client

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=ads&age=18")
	// if err != nil {
	// 	fmt.Println("get url failed,err:", err)
	// 	return
	// }

	data := url.Values{} //url value
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周")
	data.Set("age", "30")
	queryStr := data.Encode() //URL encode之后的地址
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("read resp.Body failed,err:", err)
		return
	}
	defer resp.Body.Close() //关闭resp.Body
	//从resp中把服务端返回的数据读出来
	// resp.Body.Read()
	// resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed,err:", err)
		return
	}
	fmt.Println(string(result))
}
