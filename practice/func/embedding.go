package main

import "fmt"

type Person struct{
	name string
	age int
}

func (p * Person) greeting(){
	fmt.Println("Hello~~")
}
/*
type Student struct{
	p 	Person		// has-a 관계 학생 구조체 안에 사람 구조체를 필드로 가지고 있음
	school string
	grade int 
	
}
//Go 언어는 클래스를 제공하지 않으므로 상속 또한 제공하지 않는다. 대신 Embedding을 사용하면 상속과 같은 효과를 낼 수 있다.
*/

type Student struct{
	Person // 필드명 없이 타입만 선언하면 is-a 관계가 된다.
	schoo string
	grade int
	
}

func (p * Student) greeting(){
	fmt.Println("Hello Student") //Override 부모 구조체의 메서드를 호출 하고 싶으면 s.Person.greeting() 과 같이 호출한다.
}


func main() {
	var s Student
	
	//s.p.greeting() 
	s.greeting()
	
}