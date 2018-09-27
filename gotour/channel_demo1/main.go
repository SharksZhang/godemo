package main

import (
	"fmt"
	"time"
)

//实验结论 x, y := <-c, <-c 这种方式读必须是channel中有两个值才能读到
func sum(s []int, c chan int, shouldSleep bool)  {
	if shouldSleep {
		time.Sleep(5 * time.Second)
	}
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main()  {
	arr := []int{1,3,4,5,6,7,8}
	c := make(chan int)
	go sum(arr[len(arr)/2:], c, false)
	//time.Sleep(2 * time.Second)
	go sum(arr[:(len(arr)/2)], c , false)
	x, y := <-c, <-c
	fmt.Println(x, y, x +y )

}