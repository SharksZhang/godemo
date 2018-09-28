package main

import (
	"fmt"
)


//可以使用 select 使一个channel用于生产，一个用于控制循环退出。

func main()  {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++{
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}

func fibonacci(c chan int, quit chan int)  {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y

		case <- quit:
			return
		}
	}
}

