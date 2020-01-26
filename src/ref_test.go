package main

import (
	"fmt"
	"github.com/krexspace/qgengo/core"
	"math/rand"
	"testing"
)

/*
// https://golang.org/pkg/testing/
// For recursive building
// go build ./...

// To run:
// go test
// go test -v [for verbose mode]
// Or for coverage
// go test -cover
// To run:
// go test -bench
*/

func _TestCore_1(t *testing.T) {
	var qv2 core.Qvec2
	var qv3 = core.Qvec3{234.455, 434.43, -454.9}
	var qf core.QuadFace
	fmt.Println("Test", qv2, qf, qv3, qv3.Gen3DSpacialIndexKey())
}

func _BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

const VMAX = 1000

func makeVertSlice() []core.Qvec3 {
	vlist := make([]core.Qvec3, 0, 0)
	for i := 0; i < VMAX; i++ {
		v := core.Qvec3{X: rand.Float32() * 100, Y: rand.Float32() * 100, Z: rand.Float32() * 100}
		//if i < 10 {
		//	fmt.Println(v)
		//}
		vlist = append(vlist, v)
	}
	return vlist
}

func makeVertMap() map[core.QuantizedQvec3]int {
	vlist := makeVertSlice()
	c := len(vlist)
	vmap := make(map[core.QuantizedQvec3]int)
	for i := 0; i < c; i++ {
		vmap[vlist[i].Gen3DSpacialIndexKey()] = i
	}
	return vmap
}

//>go test -v
func TestVertIndexing(t *testing.T) {
	fmt.Println("Exec: TestVertIndexing")
	vmap := makeVertMap()
	v := core.Qvec3{100.3434, 200.123433, 300.12323}
	vmap[v.Gen3DSpacialIndexKey()] = 25

	fmt.Println(vmap[v.Gen3DSpacialIndexKey()])
	t.Run("Index Lookup Test", func(t *testing.T) {
		if val := vmap[v.Gen3DSpacialIndexKey()]; val != 25 {
			t.Errorf("Error: expected %d, got %d ", 25, val)
		}
	})

	panicChecker := func() {
		r := recover()
		if r == nil {
			// No-panic is test fail
			t.Errorf("Error: Supposed to panic for large values")
		} else {
			fmt.Printf("Paniced as expected: %+v\n", r)
		}
	}
	t.Run("Large Vert Value Test", func(t *testing.T) {
		{
			defer panicChecker()
			// Using an invalid large value for vert
			v := core.Qvec3{float32(core.QVERT_INT_MAX) * 10 / core.VERT_DECIMAL_PRECISION, 200.123433, 300.12323}
			v.Gen3DSpacialIndexKey()
		}
		{
			defer panicChecker()
			// Using an invalid large value for vert
			v := core.Qvec3{1e25, 200.123433, 300.12323}
			v.Gen3DSpacialIndexKey()
		}
	})
}

// Adjust VMAX
func BenchmarkVertIndexing(b *testing.B) {
	fmt.Println("Exec: TestVertIndexing")
	vmap := makeVertMap()
	v := core.Qvec3{100.3434, 200.123433, 300.12323}
	vmap[v.Gen3DSpacialIndexKey()] = 25
	v = core.Qvec3{100.3434000001, 200.123433, 300.12323}
	fmt.Println(vmap[v.Gen3DSpacialIndexKey()])

	k := v.Gen3DSpacialIndexKey()
	b.ResetTimer()
	b.Run("Test 1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			val := vmap[k]
			val++
		}
	})
}
