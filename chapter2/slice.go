package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	var mySlice []int = myArray[1:8]

	//打印输出数组内容
	fmt.Println("数组元素 : ")
	for _, v :=range myArray {
		fmt.Print(v, " ")
	}

	//打印输出切片内容
	fmt.Println("\nslice元素 : ")
	for _,v :=range mySlice {
		fmt.Print(v," ")
	}
	fmt.Println()
}