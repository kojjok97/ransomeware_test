package main

import "fmt"


func main() {
	
	//클로저(closure)

	
	a, b := 5,10 
	
	f := func(x int) int {
		return a * x + b	// 함수선언 밖에 있는 변수 사용 가능
		
	}
	
	fmt.Println(f(10))
	
	

	
}