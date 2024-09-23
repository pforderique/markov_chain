package matrix

import (
	"fmt"
	"reflect"
	"strings"
)

type Matrix struct {
	data []float64
	dims []int
}

func (A Matrix) Size() int {
	size := 1
	for _, dim := range A.dims {
		size *= dim
	}
	return size
}

func (A Matrix) Get(coors ...int) float64 {
	if len(coors) != len(A.dims) {
		panic(fmt.Sprintf(
		`Coordinate length %d does not match matrix dimensions length of 
		%d`, len(coors), len(A.dims)))
	}

	data_idx := 0
	for idx, coor := range coors[:len(coors)-1] {
		data_idx += coor*A.dims[idx]
	}

	data_idx += coors[len(coors)-1]

	return A.data[data_idx]
}

func (A Matrix) String() string {
	// s := "["
	// // data_idx

	
	if len(A.dims) == 1 {
		return fmt.Sprint(A.data)
	}

	s := "["
	currDim := A.dims[0]
	subMatrixStrings := []string{}

	for startIdx:=0; startIdx<len(A.data); startIdx += len(A.data)/currDim {
		subData := A.data[startIdx : startIdx+len(A.data)/currDim]
		subMatrix := Matrix{subData, A.dims[1:]}
		subMatrixStrings = append(subMatrixStrings, subMatrix.String())
	}
	s += strings.Join(subMatrixStrings, "\n")
	s += "]"

	// s := "["
	// currDim := A.dims[len(A.dims)-1]
	// subMatrixStrings := []string{}
	// for startIdx:=0; startIdx<=len(A.data)-currDim; startIdx += currDim {
	// 	subData := A.data[startIdx : startIdx+currDim]
	// 	subMatrix := Matrix{subData, A.dims[:len(A.dims)-1]}
	// 	// s += subMatrix.String() + "\n"
	// 	subMatrixStrings = append(subMatrixStrings, subMatrix.String())
	// }
	// s += strings.Join(subMatrixStrings, ",\n")
	// s += ",]"

	return s
}

func (A Matrix) Add(B Matrix) Matrix {
	if !reflect.DeepEqual(A.dims, B.dims) {
		panic(fmt.Sprintf(
			"Cannot add Matrix with dimensions %v and %v", A.dims, B.dims))
	}
	sumData := make([]float64, len(A.data))
	for i := range A.data {
		sumData[i] = A.data[i] + B.data[i]
	}
	return Matrix{
		data: sumData,
		dims: A.dims,
	}
}  

// func matrixMultiplySimple(A *Matrix, B *Matrix) (C *Matrix) {
// 	return C
// }