package main

import (
	"fmt"
	"github.com/krexspace/qgengo/core"
)

func main() {
	var qv2 core.Qvec2
	var qv3 = core.Qvec3{234.455, 434.43, -454.9}
	var qf core.QuadFace
	fmt.Println("Test", qv2, qf, qv3, qv3.Gen3DSpacialIndexKey())
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
