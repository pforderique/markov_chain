package matrix

import "fmt"

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

// func (A Matrix) String() string {
// 	s := "["
// 	data_idx

// 	for _, dim := range A.dims {
// 		s += "["
// 		for idx:=0; idx<dim; idx++
// 	}
// }

func MatrixMultiply(A *Matrix, B *Matrix) (C *Matrix) {
	return C
}