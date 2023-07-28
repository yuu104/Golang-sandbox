package main

import "fmt"

func main() {
	x := 7
	y := 12

	if x >= 5 && x < 10 {
		fmt.Println("xは5以上10未満です")
	} else {
		fmt.Println("xは5以上10未満ではありません")
	}

	if y >= 5 && y < 10 {
		fmt.Println("yは5以上10未満です")
	} else {
		fmt.Println("yは5以上10未満ではありません")
	}
}