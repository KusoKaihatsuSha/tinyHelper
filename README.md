[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![godoc](https://godoc.org/github.com/KusoKaihatsuSha/tinyHelper?status.svg)](https://godoc.org/github.com/KusoKaihatsuSha/tinyHelper) [![Go Report Card](https://goreportcard.com/badge/github.com/KusoKaihatsuSha/tinyHelper?)](https://goreportcard.com/report/github.com/KusoKaihatsuSha/tinyHelper?)

# Source package with the some helping functions

> ```
> SliceRotate - Rotate slice or rotate even/odd pair
> SliceSort   - Wrapper for sorting slice
> ToNum       - Wrapper for convert into numerable types
> ToStr       - Wrapper for convert to string by 'fmt.Sprintf'
> ToBase      - Wrapper for convert to other base (2, 8, 10, 16 ...)
> ```

### Examples:

```go
package main

import (
	"fmt"

	helper "github.com/KusoKaihatsuSha/tinyHelper"
)

func main() {
	i := 0
	var inc = func() int {
		i++
		return i
	}
	mask := "%d) %[2]v - TYPE[%[2]T]\n"
	test001 := []byte("0123456789")
	fmt.Printf(mask, inc(), string(helper.SliceRotate(test001, true))) // 1
	fmt.Printf(mask, inc(), string(helper.SliceRotate(test001)))       // 2
	test002 := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	test002s := helper.SliceRotate(test002)
	helper.SliceSort(test002s)
	fmt.Printf(mask, inc(), test002s) // 3
	test003 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test003s := helper.SliceRotate(test003, true)
	helper.SliceSort(test003s)
	fmt.Printf(mask, inc(), test003s) // 4
	test004 := []rune("👺👾👺👾👾👾")
	fmt.Printf(mask, inc(), string(helper.SliceRotate(test004)))        // 5
	fmt.Printf(mask, inc(), string(helper.SliceRotate(test004, true)))  // 6
	fmt.Printf(mask, inc(), helper.ToStr("123")+".")                    // 7
	fmt.Printf(mask, inc(), helper.ToStr(123)+".")                      // 8
	fmt.Printf(mask, inc(), helper.ToStr(byte(123))+".")                // 9
	fmt.Printf(mask, inc(), helper.ToStr(float64(123))+".")             // 10
	fmt.Printf(mask, inc(), helper.ToInt("122")+1)                      // 11
	fmt.Printf(mask, inc(), helper.ToInt("122")+1)                      // 12
	fmt.Printf(mask, inc(), helper.ToNum[uint]("18446744073709551615")) // 13
	fmt.Printf(mask, inc(), helper.ToNum[int](byte(122))+1)             // 14
	fmt.Printf(mask, inc(), helper.ToNum[byte](byte(122))+1)            // 15
	fmt.Printf(mask, inc(), helper.ToNum[int](-123+1))                  // 16
	fmt.Printf(mask, inc(), helper.ToNum[float64](float64(122.1))+1)    // 17
	fmt.Printf(mask, inc(), helper.ToNum[int](true)+1)                  // 18
	fmt.Printf(mask, inc(), helper.ToNum[int](10001, 2))                // 19
	fmt.Printf(mask, inc(), helper.ToNum[int]("00010001", 2))           // 20
	fmt.Printf(mask, inc(), helper.ToNum[int]("0b010001", 2))           // 21
	fmt.Printf(mask, inc(), helper.ToNum[int]("0b10001"))               // 22
	fmt.Printf(mask, inc(), helper.ToNum[float64]("10001", 2))          // 23
	fmt.Printf(mask, inc(), helper.ToNum[int]("0o21", 8))               // 24
	fmt.Printf(mask, inc(), helper.ToNum[int]("0o21"))                  // 25
	fmt.Printf(mask, inc(), helper.ToNum[float64]("21", 8))             // 26
	fmt.Printf(mask, inc(), helper.ToNum[int]("0x11", 16))              // 27
	fmt.Printf(mask, inc(), helper.ToNum[int]("0x11"))                  // 28
	fmt.Printf(mask, inc(), helper.ToNum[float64]("11", 16))            // 29
	fmt.Printf(mask, inc(), helper.ToBase(17, 10, 2, 8))                // 30
	fmt.Printf(mask, inc(), helper.ToBase(17, 10, 8))                   // 31
	fmt.Printf(mask, inc(), helper.ToBase(17, 10, 16))                  // 32
	fmt.Printf(mask, inc(), helper.ToInt(17, 2))                        // 33 // error base, return 17
	fmt.Printf(mask, inc(), helper.ToBase(10001, 2, 16))                // 34
	fmt.Printf(mask, inc(), helper.ToBase("1a", 11, 10))                // 35
	fmt.Printf(mask, inc(), helper.ToInt("111", 36))                    // 36
	fmt.Printf(mask, inc(), helper.ToBase(30, 7, 16))                   // 37
	fmt.Printf(mask, inc(), helper.ToInt("11a", 16))                    // 38
	fmt.Printf(mask, inc(), helper.ToBase(21, 10, 11))                  // 39
	fmt.Printf(mask, inc(), helper.ToBase(21, 10, 16))                  // 40
	fmt.Printf(mask, inc(), helper.ToBase(10, 11, 16, 2))               // 41
	fmt.Printf(mask, inc(), helper.ToBase("00010001", 2, 16))           // 42
	fmt.Printf(mask, inc(), helper.ToBase("00001111", 2, 16, 2))        // 43
	fmt.Printf(mask, inc(), helper.ToBase("0o00021", 8, 16))            // 44
	fmt.Printf(mask, inc(), helper.ToNum[byte](1222))                   // 45 // error bitSize, return quotient = 1222%256
	fmt.Printf(mask, inc(), helper.ToBase("0b1", 16, 10, 2))            // 46
	fmt.Printf(mask, inc(), helper.ToBase("0b1", 2, 10, 2))             // 47
	fmt.Printf(mask, inc(), helper.ToBase(3, 10, 10, 2))                // 48
	// Output:
	// 1) 1032547698 - TYPE[string]
	// 2) 9876543210 - TYPE[string]
	// 3) [0 1 2 3 4 5 6 7 8 9] - TYPE[[]string]
	// 4) [0 1 2 3 4 5 6 7 8 9] - TYPE[[]float64]
	// 5) 👾👾👾👺👾👺 - TYPE[string]
	// 6) 👾👺👾👺👾👾 - TYPE[string]
	// 7) 123. - TYPE[string]
	// 8) 123. - TYPE[string]
	// 9) 123. - TYPE[string]
	// 10) 123. - TYPE[string]
	// 11) 123 - TYPE[int]
	// 12) 123 - TYPE[int]
	// 13) 18446744073709551615 - TYPE[uint]
	// 14) 123 - TYPE[int]
	// 15) 123 - TYPE[uint8]
	// 16) -122 - TYPE[int]
	// 17) 123.1 - TYPE[float64]
	// 18) 2 - TYPE[int]
	// 19) 17 - TYPE[int]
	// 20) 17 - TYPE[int]
	// 21) 17 - TYPE[int]
	// 22) 17 - TYPE[int]
	// 23) 17 - TYPE[float64]
	// 24) 17 - TYPE[int]
	// 25) 17 - TYPE[int]
	// 26) 17 - TYPE[float64]
	// 27) 17 - TYPE[int]
	// 28) 17 - TYPE[int]
	// 29) 17 - TYPE[float64]
	// 30) 00010001 - TYPE[string]
	// 31) 21 - TYPE[string]
	// 32) 11 - TYPE[string]
	// 33) 17 - TYPE[int]
	// 34) 11 - TYPE[string]
	// 35) 21 - TYPE[string]
	// 36) 1333 - TYPE[int]
	// 37) 15 - TYPE[string]
	// 38) 282 - TYPE[int]
	// 39) 1a - TYPE[string]
	// 40) 15 - TYPE[string]
	// 41) 0b - TYPE[string]
	// 42) 11 - TYPE[string]
	// 43) 0f - TYPE[string]
	// 44) 11 - TYPE[string]
	// 45) 198 - TYPE[uint8]
	// 46) 177 - TYPE[string]
	// 47) 01 - TYPE[string]
	// 48) 03 - TYPE[string]
}
```

### Add to use

  ```shell
$ go get github.com/KusoKaihatsuSha/tinyHelper.git
  ```
