package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int = 55

	// var i int
	// i = 55

	var f float32 = 4.30
	//f := 4.30 // float64
	i := 55

	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	//strings.Title은 유니코드 단어 경계 처리가 부정확해 표준 라이브러리에서 폐기(Deprecated)됨. 실제 코드에선 golang.org/x/text/cases 사용 권장
	fmt.Printf("%s\n", strings.Title("kim inha"))
	fmt.Println(math.Ceil(3.99))

	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n")
	fmt.Println("value i :", i)
}
