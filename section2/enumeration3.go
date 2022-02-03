//열거형3
package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println(A, B, C)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
