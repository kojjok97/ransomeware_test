package main

import "fmt"

func f() {
	
	defer func() {
		s := recover() //에러인 상황을 복구함
		fmt.Println(s) 
		
	}()
	
	a := [...]int{1,2}
	
	for i := 0; i <3 ; i++{
		fmt.Println(a[i])
	
		
	}
	
}


func main() {
	
	/*a := [...]int{1,2}
	
	for i := 0; i <3 ; i++{
		fmt.Println(a[i])
	
		에러 발생
	}*/
	
	//panic("Error !")
	fmt.Println("Hello world")
	//문법적인 에러는 아니지만 로직에 따라 에러처리 하고 싶을 경우 사용 
	
}