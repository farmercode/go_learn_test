package main

import "fmt"
/**
获得用户信息的匿名函数
 */
func GetUserInfo() (name,job string,age int)  {
	realname := "James"
	return realname, "teacher", 28
}

func main(){
	age := 26
	var name,job string
	name = "king"
	job = "programmer"
	fmt.Printf("Name:%s\tAge:%d\tJob:%s\n",name,age,job)
	name,_,_ = GetUserInfo()
	fmt.Printf("Name:%s\tAge:%d\tJob:%s\n",name,age,job)
}
