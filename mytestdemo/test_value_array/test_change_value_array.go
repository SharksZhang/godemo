package main

import "fmt"

func main()  {
	students := []Student{
		{
			Name: "s1",
			Age:  "8",
		},
		{
			Name: "s2",
			Age:  "2",
		},
	}
	//测试使用for range 获取到的值 改变是不会改变数组中真实的值
	for _, s := range students{
		s.Name ="schange"
	}
	fmt.Println(students)
	//使用索引值改变 可以改变
	for index, _:= range students{
		students[index].Name = "schange"
	}
	fmt.Println(students)

	//使用索引值赋值后改变, 无法改变数组中真实的值。是值传递
	for index, _:= range students{
		s := students[index]
		s.Name = "schange1"
	}
	fmt.Println(students)
	//使用索引值取指针后传递, 可以改变中结构体的值
	for index:= range students{
		s := &students[index]
		s.Name = "schange2"
	}
	fmt.Println(students)



}

type Student struct {
	Name string
	Age  string
}