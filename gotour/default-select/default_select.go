package main

import (
	"time"
	"fmt"
)

/**
1.time.Tick会返回一个管道，管道中每个固定时间写入当前时间
2.time.After会在固定时间后在管道中写入时间
3.记得死循环要返回。
**/
func main() {
	tick := time.Tick(100 * time.Millisecond)
	after := time.After(500 * time.Millisecond)

	for {
		select {
		case i := <-tick:
			fmt.Println("tick:[%v]", i)
		case <-after:
			fmt.Println("boom boom boom")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}

	}

}
