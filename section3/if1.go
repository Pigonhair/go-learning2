//IF문(1)

package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	//소괄호 사용 x

	var a int = 20
	b := 20

	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//에러 발생1(괄호 위치)
	/*
		if b >= 25
		{
			fmt.Println("25이상")
		}
	*/

	//에러 발생2(괄호 생략 불가능)
	/*
		if b >= 25
			fmt.Println("25이상")
	*/

	//에러 발생3(무조건 Boolean 타입이어야함)
	/*
		if c := 1; c {
			fmt.Println("True")
		}
	*/

	//예제3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}
}
