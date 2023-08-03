package main

import "fmt"

func compositeAssignmentOperator() {
	// 変数xに10を加えて、その結果をxに代入
	x := 5
	x += 10
	fmt.Println(x)

	// 変数yから3を引いて、その結果をyに代入
	y := 8
	y -= 3
	fmt.Println(y)

	// 変数zに2をかけて、その結果をzに代入
	z := 7
	z *= 2
	fmt.Println(z)

	// 変数pを2で割って、その結果をpに代入
	p := 12
	p /= 2
	fmt.Println(p)

	// 変数aに1を加算
	a := 8
	a++
	fmt.Println(a)

	// 変数bに1を減算
	b := 8
	b--
	fmt.Println(b)
}
