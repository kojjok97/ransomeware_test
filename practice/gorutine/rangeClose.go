package main

import(
	"runtime"
	"fmt"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU갯수를 구한 뒤 사용할 최대 CPU개수를 설정 
	fmt.Println(runtime.GOMAXPROCS(0))
	s := "Hello world"
	
	for i := 0 ; i < 100 ; i++ {
		go func(n int){
			fmt.Println(s,n)
		}(i)
		
	}
	
	fmt.Scanln()
	
}