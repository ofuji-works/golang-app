package main

import (
	"fmt"
)

func main() {
	// 標準出力
	fmt.Println("Hello world!")

	// 変数宣言
	number := 1
	fmt.Println(number)

	// if文
	if number != 1 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// 四則演算
	add := 1 + 2
	fmt.Println(add)

	substract := 3 - 1
	fmt.Println(substract)

	multiplication := 5 * 5
	fmt.Println(multiplication)

	division := 25 / 5
	fmt.Println(division)

	rest := 26 % 5
	fmt.Println(rest)

	// 配列
	var arr1 [2]int
	arr1[0] = 0
	arr1[1] = 1
	fmt.Println(arr1)

	var arr2 [2]int = [2]int{1, 2}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3)

	// for
	for num := 0; num < 10; num++ {
		text := fmt.Sprintf("for No.[%d]", num)
		println(text)
	}

	// switch
	switchvar := "switch"

	switch switchvar {
	case "switch":
		fmt.Println(switchvar)
		break
	default:
		fmt.Println("default")
		break
	}

}
