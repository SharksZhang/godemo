package main

import "fmt"


func main() {
	s := array{}
	fmt.Println(len(s.a))
	var b []int
	fmt.Println(len(b))
}

type array struct {
	a []int
}
