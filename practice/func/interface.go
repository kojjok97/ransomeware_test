package main


import "fmt"
import "strconv"

// 인터페이스 자체는 자료형을 가지고 있지 않는다.

type MyInt int

func(i MyInt) Print(){
	fmt.Println(i)
}

type Rectangle struct{
	width, height int
}

func(r Rectangle) Print(){
	fmt.Println(r)
}

type Printer interface{
	Print()
}


func formatString(arg interface{}) string{
	switch arg.(type){
		
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
		
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
		
	case string:
		s := arg.(string)
		return s
	default:
		return "error"
		
	}
}


func main() {

	var i MyInt = 5
	rect := Rectangle{10,20}
	
	var p Printer //인터페이스 선언 
	_ = p
	/*p = i 
	fmt.Println(p)
	p = rect
	fmt.Println(p)
	*/
	//인터페이스 선언과 동시에 초기화
	p1 := Printer(i)
	p2 := Printer(rect)
	p1.Print()
	p2.Print()
	
	pArr := []Printer{i,rect} //슬라이스 형태로 인터페이스 초기화
	
	for index,_ := range pArr{
		pArr[index].Print() //슬라이스를 순회하면서 print메서드 호출
	}
	
	for _,value := range pArr{
		value.Print() //슬라이스를 순회하면서 print메서드 호출
	}
	// 빈 인터페이스 활용 

	fmt.Println(formatString(10)) //formatString의 매개변수를 interface{}로 지정하여 모든 타입을 처리할 수 있다. 
	fmt.Println(formatString(2.5)) // 자료형 뿐만 아니라 구조체 인스턴스 및 포인터도 넘길 수 있다.
	fmt.Println(formatString("Hello world!"))
	
	
}