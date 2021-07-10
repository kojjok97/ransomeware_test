package main

import "fmt"


func main() {
	
	var numPtr1 *int //포인터형 변수를 선언하면 nil로 초기화됨
	_ = numPtr1
	
	var numPtr2 *int = new(int) //빈 포인터형은 바로 사용불가하여 new함수로 초기화해야함
	
	*numPtr2 = 1 //역참조로 값대입
	
	fmt.Println(*numPtr2)
	
}