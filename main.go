package main

import (
	"fmt"
)

// //可変長引数
// func Sum(s ...int) int {
// 	n := 0
// 	for _, v := range s {
// 		n += v
// 	}
// 	return n
// }

func main() {
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl)
	// sl := []int{1, 2, 3, 4, 5}
	// sl2 := make([]int, 5, 10)
	// fmt.Println(sl2)

	// n := copy(sl2, sl)

	// fmt.Println(n, sl2)

	// sl := []string{"A", "B", "C", "D", "E"}
	// fmt.Println(sl)

	// // for _, v := range sl {
	// // 	fmt.Println(v) //index番号を非表示
	// 	fmt.Println(sl[i])
	// }

	// fmt.Println(Sum(1, 2, 3))
	// fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}

	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)

	m4[1] = "Japan"
	m4[2] = "America"
	fmt.Println(m4)

	fmt.Println(m4[2])

	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s)

	delete(m4, 1)
	fmt.Println(m4)

}
