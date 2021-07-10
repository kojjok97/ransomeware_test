package main

import (
	"io/ioutil"
	"fmt"
)


func main(){
	
	fmt.Println("테스트용 파일")
	
	
	if fi, err := ioutil.ReadFile("./testfile.txt"); err == nil{
		fmt.Printf("%s",fi)
	}
	
	fmt.Println("for문 테스트")
	
	for i, x :=0, 0 ; i<10; i, x=i+1, x+2{
		fmt.Println(i,x)
	}
	
	fmt.Println("배열 슬라이스 테스트")
	
	arr := [...]int{1,2,3,4,5}
	
	slice := make([]int,5,10)
	
	slice = []int{1,2,3,4}
	
	for _ , v := range arr {
		
		fmt.Println(v)
		
	}
	
	for i, v := range slice{
		
		fmt.Println(i,v)
	}
	
	fmt.Println("함수 테스트")
	
	fmt.Println(sum(1,2,3,4,5,6))
	
	
	fmt.Println("map 테스트")
	
	maptest := map[string]map[string]int{
		"밥먹고싶다":map[string]int{
			"짜장면":5000,
			"치킨":18000,
			"초밥":23000,
		},
		"집가고싶다":map[string]int{
			"SRT":60000,
			"비행기":30000,
		},
		
		
	}
	
	for ke,va := range maptest{
		fmt.Println(ke)
		for k,v := range va{
			fmt.Println(k,v)
		}
		
	}
	
	maptest2 := map[string]map[string]int{
		"한식":map[string]int{
			"비빔밥":7000,
			"불고기":6000,
			"간장게장":12000,
		},
		"양식":map[string]int{
			"빵":5000,
			"스파게티":8000,
			"스테이크":20000,
		},
		
	}
	
	for k,v := range maptest2{
		fmt.Println(k)
		for i,x := range v{
			fmt.Println(i,x)
		}
		
	}
	
	var po *int = new(int)
	
	*po = 1 
	
	fmt.Println("포인터 테스트")
	fmt.Println(*po)
	fmt.Println(po)
	
}




func sum(n ... int)(int){
	
	total := 0
	
	for _, v := range n{
		
		total += v
	}
	
	return total
	
}