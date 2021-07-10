package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		
		for {
			
			fmt.Println(<-c1)
			time.Sleep(100 * time.Millisecond)
		}
		
	}()
	
	go func() {
		
		for {
			c2 <- "Hello World"
			time.Sleep(500 * time.Millisecond)
			
		}
		
	}()
	
	go func() {
		
		for {
			select{
			case c1 <- 10:
			case s := <-c2:
				fmt.Println(s)	
				
			}
			
		}
		
	}()
	
	time.Sleep(10*time.Second)
}