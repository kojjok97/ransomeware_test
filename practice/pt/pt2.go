package main

import "fmt"


type Nums struct{
	
	a, b int
	
}

func newNums(a,b int) *Nums{
	return &Nums{a,b}
	
}

func (n *Nums) sum() int{
	
	return n.a + n.b
	
}



func main(){
	
	
	
	fmt.Println("struct와 method")
	// 구조체를 생성자 처럼 만든 후 값을 더하는 메서드를 만들어서 출력 
	
	a := newNums(10,20)
	
	fmt.Println(a.sum())
	
}