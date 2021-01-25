package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//1.依次拿到所有名字
	//2.拿到一个人名根据规则分金币
	//2.1每个人分的金币数保存到map中
	//2.2还要记录剩余金币数
	for _, userName := range users {
		for _, c := range userName {
			switch c {
			case 'e', 'E':
				distribution[userName]++
				coins--
			case 'i', 'I':
				distribution[userName] += 2
				coins -= 2
			case 'o', 'O':
				distribution[userName] += 3
				coins -= 3
			case 'u', 'U':
				distribution[userName] += 4
				coins -= 4
			}
		}
	}
	// var sum int
	// for _, coin := range distribution {
	// 	sum += coin
	// }
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for key, value := range distribution {
		fmt.Println(key, value)
	}
}
