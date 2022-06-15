package main

import (
	"fmt"
)

// 構造体
type Person struct {
	age  int
	name string
}

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

	person1 := &Person{18, "tarp"}
	person2 := &Person{name: "tarp", age: 18}
	var person3 Person
	person3.name = "tarp"
	person3.age = 18
	fmt.Println(person1.name, person1.age)
	fmt.Println(person2)
	fmt.Println(person3)

	mikne := NewPerson(18, "tarp")
	fmt.Println("new person", mikne)

	// structのポインタを通してのアクセス
	person := Person{18, "tarp"}
	john := &person
	fmt.Println("person", person)
	fmt.Println("john", john)
	john.age = 55
	fmt.Println("person", person)
	fmt.Println("john", john)

	//構造体でのメソッドの定義
	//https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157
}

// Goにはコンストラクタが用意されていません。
// しかし、コンストラクタのようなものは実装できます。その時の関数名として通例、 New + 構造体名 が適用されます。
func NewPerson(age int, name string) *Person {
	// newは指定した型のポインタ型を生成する関数です。
	person := new(Person)
	person.age = age
	person.name = name
	return person
}
