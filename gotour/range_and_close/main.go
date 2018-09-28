package main

import (
	"fmt"
)


func fibonacci(n int, c chan int)  {
	x, y := 0, 1
	for i := 0; i< n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
	//如果不关闭这个channel，在range到 goroutine退出时，会panic
}


func main()  {

	c := make(chan int, 1)
	go fibonacci(11, c)
	for i := range c {
		fmt.Println(i)
	}
}

