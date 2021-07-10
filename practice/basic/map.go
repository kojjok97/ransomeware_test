package main

import "fmt"

func main(){
	
	//map 사용
	
	var i map[string]int //키는 string인 값은 int인 맵 선언 
	//map도 make를 이용해서 값을 넣을 수 있음 
	
	var j =make(map[string]int)
	k := make(map[string]int)
	l := map[string]int{"Hello": 10,"world":20}
	m := map[string]int{
		"Hello": 10,
		"world": 20, //여러줄로 나열할 때는 항상 ,를 붙인다.
		
	}
	fmt.Println(l["Hello"])
	fmt.Println(m["world"])
	_, _, _ = i,j,k
	
	solarSystem :=make(map[string]float32)
	
	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.960
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528
	
	fmt.Println(solarSystem["Earth"])
	fmt.Println(solarSystem["pluto"]) //없는 키를 조회시 0출력 자료형이 문자열이라면 ""빈 문자열 출력
	
	value, ok := solarSystem["pluto"]
	fmt.Println(value,ok)
	
	if value, ok := solarSystem["Neptune"]; ok{ // 키가 있다면 대입시 value에는 값 ok에는 bool값을 넣는다.
		fmt.Println(value)
	}
	for key, value := range solarSystem { // map 값 순회하기 반복문이 실행될 때마다 키와 값이 자동으로 변수에 들어감
		fmt.Println(key,value)
	}
	
	for _, value := range solarSystem { //마찬가지로 키값을 사용하지 않다면 _를 사용
		fmt.Println(value)
	}
	
	// 맵삭제하기 delete(맵,키) 해당 키에 해당하는 값이 삭제
	
	terrestrialPlanet := map[string]map[string]float32{ //map안에 map만들기
		"Mercury":map[string]float32{
			"meanRadius": 2439.7,
			"mass": 3.3022E+23,
			"orbitalPeriod":87.969,
		},
		"Venus":map[string]float32{
			"meanRadius": 6051.8,
			"mass": 4.8676E+24,
			"orbitalPeriod":224.70069,
		},
		
	}
	
	
	fmt.Println(terrestrialPlanet["Mercury"]["orbitalPeriod"])
	
}