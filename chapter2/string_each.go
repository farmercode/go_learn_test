//字符串遍历
package main

import "fmt"

func main() {
	//注意中文字符在utf8中占3个字节
	str1 := "abcdefg牛气冲天"
	len := len(str1)
	for i :=0; i < len; i++ {
		ch := str1[i]
		fmt.Println(i,ch)
	}
}
