package main

import "fmt"

func main(){
	
	fmt.Println(sum(5,6))
	fmt.Println(SumAndDiff(10,5))
	sum, diff :=SumAndDiff(20,10) //대입해줄 때는 두개를 순서대로 놓으면 됨
	//_, diff :=SumAndDiff(20,10) 생략 가능 
	fmt.Println(sum,diff)
	fmt.Println(sum2(1,2,3,4,5,6,7,8))
	n := []int{1,2,3,4,5,6} //슬라이스 생성
	x := sum2(n...) //가변인자 함수는 슬라이스 자체를 받을 수 없음으로 ...을 붙여서 넣어준다.
	fmt.Println(x)
	
	//함수를 변수에 저장 
	var hello func(a int, b int)  int = sum3
	world := sum3
	
	fmt.Println(hello(5,10))
	fmt.Println(world(10,20))
	
	
	
	
	//슬라이스에 함수 저장 
	f := []func(a int, b int) int{sum3,diff2}
	
	fmt.Println(f[0](10,20))
	fmt.Println(f[1](20,10))
	
	fm := map[string]func(a int, b int) int{"sum":sum3,"diff":diff2}
	
	fmt.Println(fm["sum"](30,50),fm["diff"](50,20))
	
	
	
	
	//익명 함수 쓸때 없어보이지만 go루틴,클로저,지연 호출을 사용할 때 쓰인다고 함
	func() {
		fmt.Println("Hello World")
		
	}()
	
}

func sum3 (a int, b int) int{
	return a+b
	
}

func diff2 (a int, b int) int{
	return a-b
}


func sum(a int,b int) (r int) { //int형 리턴 값을 변수 r로 사용 
	
	r = a+b
	return // return r 할 필요없이 r만 해도 리턴이 됨
	
	
}

func SumAndDiff(a int, b int) (int, int){ //리턴값을 두개이상 지정 가능 마찬가지로 변수 지정해서 리턴 가능
	
	return a+b, a-b
	
}

func sum2(n ...int) int{ //int형 가변인자를 받는 함수 정의 받을수 있는 매개변수의 갯수가 달라짐 슬라이스 타입임
	
	total := 0
	
	for _, value := range n{
		total += value
	}
	
	return total
}