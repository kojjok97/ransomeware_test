package main

import (
	"fmt"
	"os"
)


func hello() {
	fmt.Println("hello")
	
}

func world() {
	fmt.Println("world")
	
}

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() // 지연 호출한 file.Close() 가 맨 마지막에 호출됨
	
	if err != nil {
		fmt.Println("error")
		return // file.Close() 함수 호출
	}
	
	buf := make([]byte, 100)
	
	if _, err = file.Read(buf); err != nil {
		fmt.Println("err")
		return // file.Close() 함수 호출
	}
	
	fmt.Println(string(buf))
	
	// file.Close() 호출
	
}

func main() {
	
		// 지연 호출 지연 함수 
	
	defer world() //에러처리가 복잡해 질 경우 사용한다.
	hello()
	hello()  //hello 함수들이 먼저 실행되고 defer world()가 나중에 실행됨
	
	ReadHello()
	
}