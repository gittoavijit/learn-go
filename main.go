package main

import (
	"fmt"

	"appu.com/learngo/twosums"
)

// func main1() {
// 	result_string := twosums.Sum("a", "b")
// 	fmt.Println(result_string)
// 	result_int_string := twosums.Sum("a", "1")
// 	fmt.Println(result_int_string)
// 	result_int := twosums.Sum(10, 11)
// 	fmt.Println(result_int)
// 	result_float := twosums.Sum(1.1, 2.3)
// 	fmt.Println(result_float)
// }

func main2() {
	result := &twosums.Operator[int]{1, 20}
	fmt.Println(result.Sum())
	fmt.Println(result.Difference())
	result2 := &twosums.Operator[float64]{A: 1.0, B: 20.2}
	fmt.Println(result2.Sum())
	fmt.Println(result2.Difference())

	result3 := twosums.NewOperator(10, 12)
	fmt.Println(result3.SumWith(40))
}
func main() {
	main2()
}
