// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/ioutil"
	_"runtime"
) // import도 한번에 여러개 가저오는 것이 가능

func main() {
	fmt.Println("start")
	var i int //변수 선언
	var s string 
	
	i = 10 //변수 초기화
	s = "hello"
	
	age := 20 //선언과 동시에 초기화
	name := "Maria"
	
	var ( // var를 이용한 여러개의 변수 한번에 선언 하고 초기화 
		x,y  int=30,50 //변수 2개 동시에 초기화
		name2,age2 = "Dave", 20 // 변수 선언시 자료형 생략 가능
	
	)
	
	_ = x //변수 선언 후 사용하지 않을 경우 에러 발생 
	_ = y //변수 선언 후 사용하지 않을 경우 _ = 변수 지정하면 에러 발생하지 않음
	
	str :=`안녕하세요
잘가용` //여러줄인 문자를 저장할 때는 ``(백쿼드)로 묶어준다.
	
	const num int = 100 //상수 만들기 const는 생략 불가능 
	const str2 = "자료형은 생략 가능" //자료형은 생략 가능
	
	const(
		cx,cy int = 10, 20 	// 상수는 사용하지 않아도 별도의 오류는 없음
		cage,cname = 20, "David" //const 는 초기값을 반드시 지정해야함 
	)
	
	const(
		Sunday = iota //상수에 값을 일일이 대입하지 않아도 알아서 0 ~ n 까지 숫자를 대입해준다.
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		NumberOfDays
	
	)
	
	const( 
		a = 1 << iota // a == 1
		b = 1 << iota // b == 2
		c = 1 << iota // c == 4
	)
	
	const(
		d = iota * 30 // d == 0
		e = iota * 30 // e == 30
		f = iota * 30 // f == 90
	)
	const(
		bit0, mask0 = 1 << iota, 1<<iota -1 
		bit1, mask1
		_, _ // 생략하면 iota == 2 를 생략
		bit3, mask3
	
	)
	
	fmt.Println(s,i)
	fmt.Println(age,name)
	fmt.Println(name2,age2)
	fmt.Println(str)
	fmt.Println(str2)
	fmt.Println(num)
	fmt.Println(Wednesday)
	fmt.Println(NumberOfDays)
	fmt.Println(bit3,mask3)
	
	if i >=5 { // c언어 if문과 동일하지만 ()는 생략된다.
		fmt.Println(i)
		
	}else if i < 5 { //else if 도 동일하며 else문도 동일
		fmt.Println("5보다작음")
	}
	

	if fi, err := ioutil.ReadFile("./hello.txt"); err == nil{ //if문에서 함수를 실행한 후 조건을 판단
		fmt.Printf("%s", fi) // 포맷팅도 같음
	}
	
	//fmt.Println(fi) 사용불가능
	//fmt.Println(err) 사용불가능
	
	for fex := 0; fex < 10; fex++{ //for문도 동일
		fmt.Println(fex)
	}
	
	Loop: //Loop라는 레이블 지정
	for Lex := 0; Lex < 3; Lex++ { //continue도 마찬가지
		for j := 0; j <3; j++ {
			if j==2{
				break Loop // 중첩된 반복문 빠져나오기
			}
			fmt.Println(Lex,j)
		}
		
	}
	
	for g, h := 0, 0; g<10; g, h=g+1,h+2 { //이런것도 가능 대신 i++, j+=2 이런건 안됨
		fmt.Println(g,h)
	}
	
	switch i{ // case에 조건문,문자열,숫자여러개를 넣는 것이 가능
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("없음")	
	}
	
}