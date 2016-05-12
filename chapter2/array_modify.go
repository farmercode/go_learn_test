/**
需要注意数组在函数传值的过程中是值传递，而非指针，
所以函数内部的数组修改是不会影响到外部的，这个大部分编程语言类似
 */
package main

import "fmt"

func modify(array [5]int){
	array[0] = 10
	fmt.Println("In modify function,array values :" , array)
}

func main() {
	array := [5]int{1,2,3,4,5}
	modify(array)
	fmt.Println("In main function,array values :" , array)
}
