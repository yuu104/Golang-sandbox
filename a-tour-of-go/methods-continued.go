// 例で挙げたstructの型だけではなく、任意の型(type)にもメソッドを宣言できます。
// 例は、 Abs メソッドを持つ、数値型の MyFloat 型です。
// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。 他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できません （組み込みの int などの型も同様です）。

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
} 

func methodsContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
}