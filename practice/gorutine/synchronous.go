package main

import (
	"fmt"
	"time"
)


func main() {
	
	done := make(chan bool)
	count := 3
	
	go func(){
		for i := 0 ; i < count; i++{
			done <- true				 // 고루틴에 true를 보냄 값을 꺼낼 때 까지 대기
			fmt.Println("고루틴 :",i)		//반복문의 변수 출력
			time.Sleep(1*time.Second) 
			
		}
	}()
	
	for i := 0 ; i <count; i++ {
		<-done
		fmt.Println("메인 함수:",i)
	}
	
}