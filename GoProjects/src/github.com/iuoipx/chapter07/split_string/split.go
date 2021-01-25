package splitstring

import (
	"strings"
)

//切割字符串
//example:
//abc  b=>[a,c]

//Split 切割字符串
func Split(str string, sep string) []string {
	//提前申请内存,减少内存申请次数，提升性能
	//计算str中有几个sep，即计算申请多少内存
	//性能前后对比
	//BenchmarkSplit-4     2948394     395 ns/op     240 B/op     4 allocs/op
	//BenchmarkSplit-4     7321353     147 ns/op     80 B/op      1 allocs/op
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index]) //如果没有提前申请内存，会在此处申请
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
