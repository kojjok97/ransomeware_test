package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	wg := new(sync.WaitGroup)
	
	for i := 0; i < 10; i++ {
		wg.Add(1) // 반복할 때마다 대기그룹에 1씩 추가 
		go func(n int){
			fmt.Println(n)
			wg.Done() // 고루틴이 끝났다는 것을 알림 defer와 함쎄 사용할 수 있음
		}(i)
		
	}
	
	wg.Wait() //fmt.Scanln() 으로 대체했던 것을 Wait()로 
	fmt.Println("the end")
	
	
	
}