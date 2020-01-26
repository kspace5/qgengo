package core

import (
	"fmt"
	"math"
)

const VERT_DECIMAL_PRECISION = 10000

//const SPACIAL_INDEX_PRECISION = 10000

type Qvec2 struct {
	X float32
	Y float32
}

type Qvec3 struct {
	X float32
	Y float32
	Z float32
}

/*
Int precision
-----------------
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807

Use the right precision here based on bounds
*/
type qvertint int64

const QVERT_INT_MAX = math.MaxInt64

// Used as keys for maps
type QuantizedQvec3 struct {
	X qvertint
	Y qvertint
	Z qvertint
}

// Critical method for performance
// Check implementation and standard approaches
func (v *Qvec3) Gen3DSpacialIndexKey() QuantizedQvec3 {
	// TODO : Check
	vx := math.Round(float64(v.X * VERT_DECIMAL_PRECISION))
	vy := math.Round(float64(v.Y * VERT_DECIMAL_PRECISION))
	vz := math.Round(float64(v.Z * VERT_DECIMAL_PRECISION))
	fmt.Println(vx, QVERT_INT_MAX, math.Abs(vx) > QVERT_INT_MAX)
	if math.Abs(vx) > QVERT_INT_MAX ||
		math.Abs(vy) > QVERT_INT_MAX ||
		math.Abs(vz) > QVERT_INT_MAX {
		panic("Vert position value exceeded limits for gen3DSpacialIndexKey for " + fmt.Sprint(v))
	}

	return QuantizedQvec3{qvertint(vx), qvertint(vy), qvertint(vz)}
}

type QuadFace struct {
	indices    [4]int
	uvs        [4]Qvec2
	normals    [4]Qvec3
	hasUvs     bool
	hasNormals bool
}

// refer: https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
type VertGroupType int

const (
	Edge VertGroupType = iota + 1
	BorderEdge
)

type VertGroup struct {
	verts      []Qvec3
	uvScale    []float32
	hasUVScale bool
	groupType  VertGroupType
}

// Alias for threads/strings/borders of verts
type VertString VertGroup

type MeshStructure struct {
	// The vert cloud
	verts []Qvec3
	// FACE WIRING, NORMALS, UVS
	quadFaces []QuadFace
	// Current active border edge or hole for next addition iteration
	currentBorderIndices []int
	// Named vert groups or borders
	holesAndBorders map[string]VertGroup
}

type MeshStructureAux struct {
}
