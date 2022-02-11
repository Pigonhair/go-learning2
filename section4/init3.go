//go 초기화 함수(3)

package main

import (
	"fmt"

	"github.com/Pigonhair/go-learning2/section4/lib"
)

var num int32

//변수 초기화
func init() {
	num = 30
	fmt.Println("Main init start!")
}

func main() {
	fmt.Println("10 보다 큰 수? : ", lib.CheckNum(num))
}
