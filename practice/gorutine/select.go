package main

import "fmt"
import "time"


func main() {
	
	c1 := make(chan int)
	c2 := make(chan string)
	
	go func(){
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
		
	}()
	
	go func(){
		for {
			c2 <- "Hello world"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	
	go func(){
		for {
			select{
			case i := <-c1: //값을 사용하지 않는다면 <-c1: 처럼 생략이 가능하다.
				fmt.Println("c1 :", i)
			case s := <-c2:
				fmt.Println("c2 :", s)
				
			}
			
		}
	}()
	
	time.Sleep(10 * time.Second)
	
}