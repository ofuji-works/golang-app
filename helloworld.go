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
	add = 4 + 5
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

	// defer 遅延実行（呼び出し元の関数がreturnされた時にLIFOで呼ばれる）
	defer fmt.Println("go !!")
	fmt.Println("Let's")

	// ポインタの取得
	pointer1 := 10
	// *int → ポインタ型
	// &<変数>でポインタを表す 参照
	// ポインタ型変数 メモリのアドレスを値として格納。ポインター型変数にしたい変数の方に*をつける
	var pointer2 *int = &pointer1
	fmt.Println(pointer1)
	// ポインタの出力
	fmt.Println(pointer2)
	// 値の参照 ポインタ変数に*をつけて出力すると値の参照ができる。
	fmt.Println(*pointer2)

}
