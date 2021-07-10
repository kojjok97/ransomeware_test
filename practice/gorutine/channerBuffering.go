package main

import (
	"fmt"
	"runtime"
	"time"
)


func main() {
	
	runtime.GOMAXPROCS(1)
	
	done := make(chan bool, 2) //버퍼가 1개이상일 경우 비동기 채널 생성
	count := 4
	
	go func() {
		for i := 0; i < count ; i++{
			done <- true
			fmt.Println("고루틴 :",i)
			time.Sleep(1*time.Second)
			
		}
	}()
	
	for i := 0 ; i < count ; i++ {
		<- done
		fmt.Println("메인함수 :" ,i)
		time.Sleep(1*time.Second)
		
	}
	
}