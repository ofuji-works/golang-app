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

	//　スライス
	var slice []string
	fmt.Println(slice)
	slice2 := []string{"1", "2"}
	fmt.Println(slice2)

	// スライス（または配列）の中から値を取り出し、新たなスライスを作成
	slice3 := slice2[0:2]
	fmt.Println(slice3)

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

	person1 := &Person{Age: 18, Name: "struct"}
	var person3 Person
	person3.Name = "tarp"
	person3.Age = 18
	fmt.Println(person1.Name, person1.Age)
	fmt.Println(person3)

	mikne := NewPerson(18, "tarp")
	fmt.Println("new person", mikne)

	// structのポインタを通してのアクセス
	person := Person{Age: 18, Name: "tarp"}
	john := &person
	fmt.Println("person", person)
	fmt.Println("john", john)
	john.Age = 55
	fmt.Println("person", person)
	fmt.Println("john", john)
	greeting := person.Greeting()
	fmt.Println(greeting)

	superPerson := new(SuperPerson)
	superPerson.Age = 88
	superPerson.Name = "taro"
	fmt.Println(superPerson.Age, superPerson.Name+"Hello ")
}

type Person struct {
	// 頭文字大文字のプロパティは外部からアクセスできる。
	Age  int
	Name string
	// 頭文字大文字のプロパティは外部からアクセスできない。
	private_name string
}

// 構造体の埋め込み
type SuperPerson struct {
	Person
}

// 関数を大文字から始めることで外部に公開される
func (p Person) Greeting() string {
	fmt.Println(p.Name + "Hello!!")
	return p.Name + "Hello!!"
}

// 構造体でのメソッドの定義
// レシーバ 引数に構造体を受け取る
func (p Person) greeting() string {
	return p.Name + "Hello!"
}

// Goにはコンストラクタが用意されていません。
// しかし、コンストラクタのようなものは実装できます。その時の関数名として通例、 New + 構造体名 が適用されます。
func NewPerson(age int, name string) *Person {
	// newは指定した型のポインタ型を生成する関数です。
	person := new(Person)
	person.Age = age
	person.Name = name
	return person
}
