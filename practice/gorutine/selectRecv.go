package main

import "fmt"

func producer(c chan<- int) { // 보내기 전용 채널 
	for i:= 0; i < 5 ; i++{
		c <- i
	}
	
	c <- 100
	
	// fmt.Println(<-c)
	
}


func consumer(c <-chan int){ // 받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}
	
	fmt.Println(<-c) //값을 꺼냄
	
	// c <- 1 채널에 값을 보내면 컴파일 에러 
}

func num(a,b int) <-chan int {
	out := make(chan int)
	go func(){
		out <- a
		out <- b
		close(out)
	}()
	
	return out
	
}

func sum(c <- chan int) <-chan int{
	out := make(chan int)
	go func() {
		r := 0
		for i := range c{
			r = r+i
		}
		out <- r
		
	}()
	
	return out
	
}

func main() {
	c := make(chan int)
	
	go producer(c)
	go consumer(c)
	
	fmt.Scanln()
	
	b := num(1,2)
	out := sum(b)
	
	fmt.Println(<-out)
	
	fmt.Scanln()
	
}