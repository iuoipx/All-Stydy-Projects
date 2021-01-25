package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件内容

func writeFile1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
		return
	}
	//write
	fileObj.Write([]byte("zhou mengbi!\n"))
	//writeString
	fileObj.WriteString("zhoulinjieshibuliao!")
	fileObj.Close()
}

func writeFile2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hellosad\n") //写到缓存中
	wr.Flush()                   //将缓存中的内容写到文件中
}

func writeFile3() {
	str := "沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed,err:%v", err)
		return
	}
}

func main() {
	writeFile2()
	//writeFile3()
}
