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

func (A Matrix) stride(dim int) int {
    stride := 1
    for i := dim + 1; i < len(A.dims); i++ {
        stride *= A.dims[i]
    }
    return stride
}

func (A Matrix) Get(coors ...int) float64 {
    if len(coors) != len(A.dims) {
        panic(fmt.Sprintf(
            `Coordinate length %d does not match matrix dimensions length of 
            %d`, len(coors), len(A.dims)))
    }

    index := 0
    for i, dim := range A.dims {
        if coors[i] >= dim || coors[i] < 0 {
            panic(fmt.Sprintf("Index %d out of range for dimension %d", coors[i], i))
        }
        index += coors[i] * A.stride(i)
    }

    return A.data[index]
}

func (A *Matrix) Set(coors []int, value float64) {
    if len(coors) != len(A.dims) {
        panic(fmt.Sprintf(
            "Coor length %d != dimensions length %d", len(coors), len(A.dims)))
    }

    index := 0
    for i, dim := range A.dims {
        if coors[i] >= dim || coors[i] < 0 {
            panic(fmt.Sprintf("Index %d out of range for dimension %d", coors[i], i))
        }
        index += coors[i] * A.stride(i)
    }

    A.data[index] = value
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

func (A *Matrix) Add(B *Matrix) *Matrix {
	if !reflect.DeepEqual(A.dims, B.dims) {
		panic(fmt.Sprintf(
			"Cannot add Matrix with dimensions %v and %v", A.dims, B.dims))
	}
	for i := range B.data {
		A.data[i] = A.data[i] + B.data[i]
	}
	return A
}

func (A *Matrix) Multiply(B *Matrix) *Matrix {
	if len(A.dims) != len(B.dims) || len(B.dims) != 2 {
		panic("Both matrices must be 2D")
	}
	if A.dims[1] != B.dims[0] {
		panic(fmt.Sprintf(
			"Matrix dimensions %v x %v cannot be multiplied", A.dims, B.dims))
	}
	return matrixMultiplySimple(A, B)
}

func matrixMultiplySimple(A *Matrix, B *Matrix) *Matrix {
	I, J, K := A.dims[0], A.dims[1], B.dims[1]

	C := Matrix{make([]float64, I*K), []int{I, K}}
	for i:=0; i<I; i++ {
		for k:=0; k<K; k++ {
			var value float64 = 0;
			for j:=0; j<J; j++ {
				value += A.Get(i, j) * B.Get(j, k)
			}
			C.Set([]int{i, k}, value)
		}
	}
	return &C
}