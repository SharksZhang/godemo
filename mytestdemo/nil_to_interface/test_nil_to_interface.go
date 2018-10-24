package main

import "fmt"

//测试：nil可不可以转换为指定对象，
//结论：不可以转换，会panic

func main()  {
	var r reader

	rr, ok := r.(*testReader)
	if ok {
		fmt.Println("rr:[%v]", rr)
	}
}

type reader interface {
	read()
}

type testReader struct {

}

func (*testReader) read() {
	fmt.Print("aaaa")
}


