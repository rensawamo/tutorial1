package main

import (
	"fmt"
	"strconv"
)

// Go言語ではプログラムのエントリポイントを作成
func main() {
	firstName := "John" // ここで確保したメモリを　updateName の関数でいじる
	// sum := sum(os.Args[1], os.Args[2])
	// fmt.Println("Sum:", sum)
	updateName(&firstName) // firstNameを参照しにいく
	fmt.Println(firstName)
}

// 関数を作成したらmain (entry point) で呼び出す
func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

// ポインタでメモリをいじる
func updateName(name *string) {
	*name = "David"
}
