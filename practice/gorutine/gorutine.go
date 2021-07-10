package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int){
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
	
}

func main(){
	for i := 0 ; i < 100 ; i++{
		go hello(i)
	}
	
	fmt.Scanln() //main함수가 종료되지 않도록 대기
	
}