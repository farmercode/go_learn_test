//字符串遍历
package main

import "fmt"

func main() {
	str1 := "abcdefg"
	len := len(str1)
	for i :=0; i < len; i++ {
		ch := str1[i]
		fmt.Println(i,ch)
	}
}
