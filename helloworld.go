package main

import ("fmt");

func main () {
	// 標準出力
	fmt.Println("Hello world!")

	// 変数宣言
	number := 1
	fmt.Println(number)

	// if文
	if (number != 1) {
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
}