package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jpy"
)

func main() {
	// s1, _ := jconsole.Input("start: ")
	// start, _ := strconv.Atoi(s1)
	// s2, _ := jconsole.Input("end: ")
	// end, _ := strconv.Atoi(s2)
	// s3, _ := jconsole.Input("step (default=1): ")
	// var step int
	// if len(s3) == 0 {
	// 	step = 1
	// } else {
	// 	step, _ = strconv.Atoi(s3)
	// }
	// result := jpy.PyRange(start, end, step)
	// fmt.Println(result)

	fmt.Println(jpy.PyRange(0, 5))
	fmt.Println(jpy.PyRange(3, 7))
	fmt.Println(jpy.PyRange(3, 4))
	fmt.Println(jpy.PyRange(3, 3))
	fmt.Println(jpy.PyRange(10))
	fmt.Println(jpy.PyRange(1))
	fmt.Println(jpy.PyRange(0))
	fmt.Println(jpy.PyRange(-4))
	fmt.Println(jpy.PyRange(4, 20, 2))
	fmt.Println(jpy.PyRange(4, 10, 1))
	fmt.Println(jpy.PyRange(10, 4, 1))

	fmt.Println(jpy.PyRange(10, 0, -1))
	fmt.Println(jpy.PyRange(35, 25, -2))
}
