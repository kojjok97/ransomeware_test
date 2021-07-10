package main


import "fmt"


type Duck struct{
}

type Person struct{
}

func (d Duck) Quack(){
	fmt.Println("꽥")
	
}

func (p Person) Quack(){
	fmt.Println("꽉")
}

func (d Duck) feather(){
	fmt.Println("깃털이 있음")
}

func (p Person) feather(){
	fmt.Println("깃털을 들고 있음")
}
type Quacker interface{
	Quack()
	feather()
}

func intheforest(q Quacker){
	q.Quack()
	q.feather()
}



func main(){
	var donald Duck
	var jonh Person
	
	intheforest(donald)
	intheforest(jonh)
	
}