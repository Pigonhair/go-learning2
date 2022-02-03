//열거형2
package main

import "fmt"

func main() {
	const (
		A = iota * 10 //0부터 시작해서 1씩증가
		B
		C
	)

	fmt.Println(A, B, C)
}
