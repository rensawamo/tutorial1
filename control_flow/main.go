package main

import (
	"fmt"
	"strconv"
)

// go の条件分
// func main() {
// 	x := 27
// 	fmt.Println(x)
// 	if x%2 == 0 {
// 		fmt.Println(x, "is even")
// 	}
// }

// 複合 if ステートメント
// if で変数をつくって　条件分岐をかけられる
func givemeanumber() int {
	return -1
}

// func main() {
// 	if num := givemeanumber(); num < 0 {
// 		fmt.Println(num, "is negative")
// 	} else if num < 10 {
// 		fmt.Println(num, "has only one digit")
// 	} else {
// 		fmt.Println(num, "has multiple digits")
// 	}
// }

// switch ステートメントを使用
// func main() {
// 	sec := time.Now().Unix()
// 	rand.Seed(sec)
// 	i := rand.Int31n(10)

// 	switch i {
// 	case 0:
// 		fmt.Print("zero...")
// 	case 1:
// 		fmt.Print("one...")
// 	case 2:
// 		fmt.Print("two...")
// 	default:
// 		fmt.Print("no match...")
// 	}
// 	fmt.Println("ok")
// }

// swifch に条件分を含めるタイプ

// func main() {
// 	switch num := 15; { // 条件式をかいて
// 	case num < 50: // case で分岐
// 		fmt.Printf("%d is less than 50\n", num)
// 		fallthrough
// 	case num > 100:
// 		fmt.Printf("%d is greater than 100\n", num)
// 		fallthrough
// 	case num < 200:
// 		fmt.Printf("%d is less than 200", num)
// 	}
// }

// defer で　遅れさせる  // ファイルをとじるときとか
// func main() {
// 	for i := 1; i <= 4; i++ {
// 		defer fmt.Println("deferred", -i)
// 		fmt.Println("regular", i)
// 	}
// }

// io は　ファイルの読み書き、デーたストリームの読み書きと基本的なインターフェイスがふくまれています
// func main() {
// 	newfile, error := os.Create("learnGo.txt")

// 	if error != nil {
// 		fmt.Println("Error: Could not create file.")
// 		return
// 	}
// 	defer newfile.Close() // lerno.txt を作って最後にとじる

// 	if _, error = io.WriteString(newfile, "Learning Go!  aaa"); error != nil {
// 		fmt.Println("Error: Could not write to file.")
// 		return
// 	}

// 	newfile.Sync()
// }

// panic 関数　：　範囲外のインデックスを使用して配列にアクセスしようとしたり、nil ポインターを逆参照したりして、実行時エラーが発生すると、Go プログラムはパニック状態になる
// プログラムが直ちに停止してコンソールに出力

// リカバー関数　：　パニックをキャッチしてプログラムの実行を継続させるために使用

//  fizbaz でプログラムをくむ
// strconv.Itoa はGoの標準ライブラリで　　int 型を　String に変える

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return "no match"
		// 		fmt.Print("no match...")
	}
	return strconv.Itoa(num)
}
func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(fizzbuzz(num)) // int → string を　printlnで出力
	}
}
