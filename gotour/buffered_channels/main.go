package main

import "fmt"

//多一个从管道中读会死锁
func main()  {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println( <- c)
	fmt.Println( <- c)

}
