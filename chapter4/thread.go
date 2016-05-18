package main

import (
	"fmt"
	"sync"
	"runtime"
)

var counter int = 0

/**
count打印函数
 */
func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main()  {
	lock := &sync.Mutex{}

	//并发调用count函数
	for i :=0; i< 10 ; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
