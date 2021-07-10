package main

import "fmt"

type Rectangle struct {
	
	width int
	height int
	
}

func (rect *Rectangle) area() int{  //(rect* Rectangle) 리시버 정의 부분 리시버는 C++ java에서 this와 비슷
	return rect.width * rect.height
	
}
func (rect *Rectangle) ScaleA(factor int){  //(rect* Rectangle) 리시버 정의 부분 리시버는 C++ java에서 this와 비슷
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}


func (rect Rectangle) ScaleB(factor int){ //  ( _ Rectangle) _로 리시버를 생략가능 
	rect.width = rect.width * factor
	rect.height = rect.height * factor
	
}


func main() {
	
	rect := Rectangle{10,20}
	
	fmt.Println(rect.area())
	
	rect1 := Rectangle{30,30}
	rect1.ScaleA(10)
	fmt.Println(rect1)
	
	rect2 := Rectangle{30,30}
	rect2.ScaleB(10)
	fmt.Println(rect2)
	
	
	
}