//패키지 접근제어(2)

package main

import (
	"fmt"

	checkUp "github.com/Pigonhair/go-learning2/section4/lib"
	_ "github.com/Pigonhair/go-learning2/section4/lib2" //빈식별자(_) : import해놓고 사용하지 않을 때
)

func main() {
	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNum(11))

}
