package tinyHelper

import (
	"fmt"
)

func Example_func() {
	test001 := []byte("0123456789")
	fmt.Println(string(SliceRotate(test001, true)))
	fmt.Println(string(SliceRotate(test001)))
	test002 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	test002s := SliceRotate(test002)
	SliceSort(test002s)
	fmt.Println(test002s)
	test003 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test003s := SliceRotate(test003, true)
	SliceSort(test003s)
	fmt.Println(test003s)
	test004 := []rune("👺👾👺👾👾👾")
	fmt.Println(string(SliceRotate(test004)))
	fmt.Println(string(SliceRotate(test004, true)))
	fmt.Println(ToStr("123") + ".")
	fmt.Println(ToStr(123) + ".")
	fmt.Println(ToStr(byte(123)) + ".")
	fmt.Println(ToStr(float64(123)) + ".")
	fmt.Println(ToInt("122") + 1)
	fmt.Println(ToNum[uint]("18446744073709551615"))
	fmt.Println(ToNum[int](byte(122)) + 1)
	fmt.Println(ToNum[float64](float64(122.1)) + 1)
	fmt.Println(ToNum[int](true) + 1)

	// Output:
	// 1032547698
	// 9876543210
	// [0 1 2 3 4 5 6 7 8 9]
	// [0 1 2 3 4 5 6 7 8 9]
	// 👾👾👾👺👾👺
	// 👾👺👾👺👾👾
	// 123.
	// 123.
	// 123.
	// 123.
	// 123
	// 18446744073709551615
	// 123
	// 123.1
	// 2
}
