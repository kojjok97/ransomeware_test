package main

import "fmt"

func main(){
	//배열 
	
	var a [5]int // int형 이며 길이가 5인 배열
	a[2] = 7 //배열 세번째에 7 대입 
	fmt.Println(a)
	
	b :=[5]int{32,29,10,200,40} // 배열 생성하면서 초기값 대입
	
	c := [5]int{
		32,
		29,
		100,
		500,
		400,
	} //여러줄로 만들 때는 요소 마지막에 ,를 붙인다.
	

	
	d :=[...]int{100,200,300,400,500} //초기화할 요소는 5개이며 ...을 사용했으므로 배열의 크기는 5가 된다.

	for i, value := range b { //배열을 대입하면 인덱스와 값이 같이 대입됨 
		fmt.Println(i,value)
	}
	
	for _, value := range d { // _으로 인덱스값 대입을 생략 가능
		fmt.Println(value)
	}
	
	d = c //배열 전체가 복사되게 됨
	
	fmt.Println(d) 
	fmt.Println(c)
	// 슬라이스 
	// 배열과 달리 []안에 길이를 지정하지 않는다. 이렇게 생성된 슬라이스는 길이가 0
	// make함수를 통해서 공간을 할당해야 사용 가능 
	
	var e []int = make([]int,5) //길이가 5인 int형 슬라이스 생성 
	var f =make([]int,10) // 자료형이 생략 가능
	g := make([]int,5) //길이가 5인 int형 슬라이스 생성 
	
	h := make([]int,5,10)//길이가 5이며 용량이 10인 슬라이스 생성 
	_, _, _ = f,g,h
	//cap으로 용량을 len으로 길이를 구할 수 있음
	//용량이 길이보다 크더라도 용량을 벗어난 인덱스에는 접근이 불가능하다
	//append함수를 사용하면 슬라이스 맨 뒤에 값을 추가 가능 
	e = []int{1,2,3}
	fmt.Println(e)
	e = append(e,4,5,6)
	fmt.Println(e)
	
	//슬라이스는 레퍼런스 타입 (배열에 대한 포인터)
	//슬라이스는 e = f 이런식으로 대입하면 주소만 복사됨으로 copy함수를 사용해야함 copy(e,f) f의 요소를 e에 복사
	fmt.Println(e[0:6])
	fmt.Println(e[1:6])
	fmt.Println(e[3:5])
	

	
}