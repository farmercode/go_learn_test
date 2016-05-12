package main

import "fmt"

func main()  {
	mySlice := make([]int,5,20)
	mySlice1 := []int{11,12,13}

	fmt.Println("mySlice : ", mySlice)
	fmt.Println("mySlice1 : ", mySlice1)
	mySlice = append(mySlice,1,2,3)
	fmt.Println("mySlice : ", mySlice)
	mySlice = append(mySlice,mySlice1...)
	fmt.Println("mySlice : ", mySlice)
	//复制mySlice1中的所有元素到mySlice中对应的位置
	copy(mySlice,mySlice1)
	fmt.Println("mySlice : ", mySlice)

	fmt.Println("len(mySlice) : " ,len(mySlice))
	fmt.Println("cap(mySlice) : " ,cap(mySlice))
}
