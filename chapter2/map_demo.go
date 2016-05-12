/**
map使用demo
 */
package main

import "fmt"

type PersonInfo struct {
	ID	string
	Name	string
	Address	string
}

func main() {
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	//db插入数据
	personDB["12345"] = PersonInfo{"001","King","China"}
	personDB["123"] = PersonInfo{"002","Tom","USA"}

	person, ok := personDB["12345"]

	if ok {
		fmt.Println("Found person",person.Name,"with ID 12345.")
	}else {
		fmt.Println("Did not find person with ID 12345.")
	}
	//删除map
	delete(personDB,"12345")

	person1, ok1 := personDB["12345"]

	if ok1 {
		fmt.Println("Found person",person1.Name,"with ID 12345.")
	}else {
		fmt.Println("Did not find person with ID 12345.")
	}
}
