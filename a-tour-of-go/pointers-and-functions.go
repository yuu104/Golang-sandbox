// ここで、 Abs と Scale メソッドは関数として書きなおしてあります。

// 再度、line 16から * を消してください。 なぜ振る舞いが変わったのかわかりますか？ コンパイルするために、さらに何が必要でしょうか。

// (よくわからなくても、次のページに行きましょう)

package main

import (
	"fmt"
	"math"
)

type VertexPointersAndFunc struct {
	X, Y float64
}

func Abs(v VertexPointersAndFunc) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *VertexPointersAndFunc, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerAndFunc() {
	v := VertexPointersAndFunc{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}