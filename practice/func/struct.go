package main

import "fmt"

type Rectangle struct{
	width int
	height int
	// width height int 같은 자료형은 한줄로 나열해도됨
}
func NewRectangle(width,height int) *Rectangle{
	return &Rectangle{width,height}
	
}

func rectangleArea(rect * Rectangle) int {
	return rect.width * rect.height
	
}

func main() {
	
	//var rect Rectangle 변수 선언 
	
	var rect1 *Rectangle 
	
	rect1 = new(Rectangle)
	rect2 := new(Rectangle)
	//구조체 인스턴스 만들면서 값 대입
	var rect3 Rectangle = Rectangle{10,20}
	rect4 := Rectangle{20,30}
	rect5 := Rectangle{width : 20,height: 30}
	
	rect1.height = 30
	rect2.width = 62 // 구조체 포인터에 접근할 때나 인스턴스 접근할 때는 .으로 통일됨 
	
	_,_,_ = rect2,rect4,rect5
	
	
	
	fmt.Println(rect1)
	fmt.Println(rect3)
	
	// 메모리 할당과 동시에 값을 초기화 하는 방법은 없지만 생성자를 흉내내는 것은 가능 
	rect6 := NewRectangle(20,30)
	
	fmt.Println(rect6)
	// 짧게 하면 이렇게 됨 
	rect7 := &Rectangle{30,40}
	
	fmt.Println(rect7)
	
	rect8 := Rectangle{30,30}
	area := rectangleArea(&rect8)
	
	fmt.Println(area)
	
}