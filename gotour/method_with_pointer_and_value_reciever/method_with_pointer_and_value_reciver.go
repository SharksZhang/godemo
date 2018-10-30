package main

//使用值接受者声明的方法不会改变调用者的值，使用指针接受者调用的方法会改变调用者的值。
import (
"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale_Pointer_reciever(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale_value_reciever(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main() {
	v_pointer := &Vertex{3, 4}
	fmt.Printf("Before v_pointer.Scale_Pointer_reciever: %+v_pointer\n", v_pointer)
	v_pointer.Scale_Pointer_reciever(5)
	fmt.Printf("After v_pointer.Scale_Pointer_reciever: %+v_pointer\n", v_pointer)
	fmt.Println()

	v_value := Vertex{3, 4}
	fmt.Printf("Before v_value.Scale_Pointer_reciever: %+v_pointer\n", v_value)
	v_value.Scale_Pointer_reciever(5)
	fmt.Printf("After v_value.Scale_Pointer_reciever: %+v_pointer\n", v_value)

	fmt.Printf("call Scale_value_reciever\n")
	fmt.Println()
	v1_pointer := &Vertex{3, 4}
	fmt.Printf("Before v1_value.Scale_value_reciever: %+v_pointer\n", v1_pointer)
	v1_pointer.Scale_value_reciever(5)
	fmt.Printf("After v1_value.Scale_value_reciever: %+v_pointer\n", v1_pointer)

	fmt.Println()
	v1_value := Vertex{3, 4}
	fmt.Printf("Before v1_value.Scale_value_reciever: %+v_pointer\n", v1_pointer)
	v1_value.Scale_value_reciever(5)
	fmt.Printf("After v1_value.Scale_value_reciever: %+v_pointer\n", v1_pointer)

}
