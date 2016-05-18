/**
代码会没有没有打印任何信息退出
 */
package main

import "fmt"

func Add(x,y int) {
	z := x + y
	fmt.Println(z)
}

func main()  {
	for i :=0; i< 10 ; i++ {
		go Add(i,i)
	}
}
