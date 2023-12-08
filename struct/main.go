package main

import (
	"fmt"
)

// 配列を宣言してみる
// 初期値はメモリに０を割り当てられている
// func main() {
// 	var a [3]int
// 	a[1] = 10
// 	fmt.Println(a[0])
// 	fmt.Println(a[1])
// 	fmt.Println(a[len(a)-1])
// }

// 配列を初期化する
// [max]型｛初期値｝
// func main() {
// 	cities := [5]string{"NY", "Paries", "Berlin", "Madrid"}
// 	fmt.Println("cities", cities)
// }

// 省略記号をもちいて最後の文字を指定できる
// keyのような感じで指定できるのか
// func main() {
// 	numbers := [...]int{99: -1}
// 	fmt.Println("First Position:", numbers[0])
// 	fmt.Println("Last Position:", numbers[99])
// 	fmt.Println("Length:", len(numbers))
// }

// 多次元配列
// func main() {
// 	var twoD [3][5]int
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 5; j++ {
// 			twoD[i][j] = (i + 1) * (j + 1)
// 		}
// 		fmt.Println("Row", i, twoD[i])
// 	}
// 	fmt.Println("\nAll at once:", twoD)
// }

// スライスについて　学習する
// func main() {
// 	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
// 	fmt.Println(months)
// 	fmt.Println("Length:", len(months))   // 要素数
// 	fmt.Println("Capacity:", cap(months)) // キャパシティ
// }

// 注意ポイント
// capa が自動拡張されるから、メモリの増大にかなり注意が必要

// func main() {
// 	var numbers []int
// 	for i := 0; i < 10; i++ {
// 		numbers = append(numbers, i)
// 		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
// 	}
// }

// マップを使用する
// func main() {
// 	studentAge := map[string]int{
// 		"John": 32,
// 		"bob":  31,
// 	}
// 	fmt.Println("Christy", studentAge["John"]) // 辞書的な感じでとれる
// 	delete(studentAge, "John")                 // deleteで　消せる
// 	fmt.Println(studentAge)
// }

// map 内でループできる
// keyとvalueがとれる
// func main() {
// 	studentsAge := make(map[string]int)
// 	studentsAge["john"] = 32
// 	studentsAge["bob"] = 31
// 	for name, age := range studentsAge {
// 		fmt.Printf("%s\t%d\n", name, age)
// 	}
// }

// 構造体を定義
// 別の構造体の埋め込みも可能
// type Employee struct {
// 	ID        int
// 	FirstName string
// 	LastName  string
// 	Address   string
// }

// func main() {
// 	employee := Employee{LastName: "Doe", FirstName: "John"}
// 	fmt.Println(employee)
// 	employeeCopy := &employee // 構造体のインスタンスはここでアドレスを指定してあげることでメモリへの参照が可能
// 	employeeCopy.FirstName = "David"
// 	fmt.Println(employee)
// }

// type Person struct {
// 	ID        int
// 	FirstName string `json:"name"`
// 	LastName  string
// 	Address   string `json:"address,omitempty"`
// }

// type Employee struct {
// 	Person
// 	ManagerID int
// }

// type Contractor struct {
// 	Person
// 	CompanyID int
// }

// func main() {
// 	employees := []Employee{
// 		Employee{
// 			Person: Person{
// 				LastName: "Doe", FirstName: "John",
// 			},
// 		},
// 		Employee{
// 			Person: Person{
// 				LastName: "Campbell", FirstName: "David",
// 			},
// 		},
// 	}
// 	//構造体を使うのもswiftとおなじ
// 	// to json jsonの文字列にマーシャリングする
// 	data, _ := json.Marshal(employees)
// 	fmt.Printf("%s\n", data)

// 	// from json
// 	var decoded []Employee
// 	json.Unmarshal(data, &decoded)
// 	fmt.Printf("%v", decoded)
// }

// データ型

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{ // 組み込みでローマ数字に対応する英数字が返される
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// mapでからの配列にappend していく感じで要素集合をつくっていく感じ
	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		//  mapに相当するキーがあるかどうか
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	total := 0

	fmt.Println(arabicVals)
	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}

func main() {
	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}
