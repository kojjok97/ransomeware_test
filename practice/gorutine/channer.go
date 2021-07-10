package main

import (
	"fmt"
)

func sum(a,b int, c chan int){
	c <- a +b
	
}

func main() {
	c := make(chan int) //채널을 사용하기 전에는 make로 무조건 공간을 할당해 줘야함
	
	go sum(1, 2, c)
	
	n := <-c
	fmt.Println(n)
	
}