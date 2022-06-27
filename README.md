[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)[![godoc](https://godoc.org/github.com/KusoKaihatsuSha/tinyHelper?status.svg)](https://godoc.org/github.com/KusoKaihatsuSha/tinyHelper)[![Go Report](https://goreportcard.com/badge/github.com/KusoKaihatsuSha/tinyHelper)](https://goreportcard.com/report/github.com/KusoKaihatsuSha/tinyHelper)

# Source package with the some helping functions

> ```
> SliceRotate - Rotate slice or rotate even/odd pair
> SliceSort   - Wrapper for sorting slice
> ToNum       - Wrapper for convert into numerable types
> ToStr       - Wrapper for convert to string by 'fmt.Sprintf'
> ```

### Examples:

```go
package main

import (
	"fmt"
	helper "github.com/KusoKaihatsuSha/tinyHelper"
)

func main() {
	test001 := []byte("0123456789")
	fmt.Println(string(helper.SliceRotate(test001, true)))
	fmt.Println(string(helper.SliceRotate(test001)))
	test002 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	test002s := helper.SliceRotate(test002)
	helper.SliceSort(test002s)
	fmt.Println(test002s)
	test003 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test003s := helper.SliceRotate(test003, true)
	helper.SliceSort(test003s)
	fmt.Println(test003s)
	test004 := []rune("👺👾👺👾👾👾")
	fmt.Println(string(helper.SliceRotate(test004)))
	fmt.Println(string(helper.SliceRotate(test004, true)))
	fmt.Println(helper.ToStr("123") + ".")
	fmt.Println(helper.ToStr(123) + ".")
	fmt.Println(helper.ToStr(byte(123)) + ".")
	fmt.Println(helper.ToStr(float64(123)) + ".")
	fmt.Println(helper.ToInt("122") + 1)
	fmt.Println(helper.ToNum[uint]("18446744073709551615"))
	fmt.Println(helper.ToNum[int](byte(122)) + 1)
	fmt.Println(helper.ToNum[float64](float64(122.1)) + 1)
	fmt.Println(helper.ToNum[int](true) + 1)

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
```

### Add to use

  ```shell
$ go get github.com/KusoKaihatsuSha/tinyHelper.git
  ```
