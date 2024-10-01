package linalg

import (
	"math/rand"
	"reflect"
	"testing"
)

const (
	zero = iota
	ones
	identity
	random
)

func createSquareMatrix(init, n int) *SquareMatrix {
	var data []float64
	switch init {
	case zero:
		data = make([]float64, n*n)
	case ones:
		data = make([]float64, n*n)
		for i := 0; i < n*n; i++ {
			data[i] = 1
		}
	case identity:
		data = make([]float64, n*n)
		for i := 0; i < n; i++ {
			data[i*n+i] = 1
		}
	case random:
		data = make([]float64, n*n)
		for i := 0; i < n*n; i++ {
			data[i] = rand.Float64()
		}
	}
	return &SquareMatrix{data, n}	
}

// TestSize calls SquareMatrix's Size() module with a 3x3 SquareMatrix. !
func TestSqSize(t *testing.T) {
	t.Parallel()
    A := SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3}
	size := A.Size()

	if size != 9 {
		t.Fatalf(`3x3 SquareMatrix A has Size() of %d instead of 9`, size)
	}
}

// TestGet calls SquareMatrix.Get
func TestSqGetSuccess(t *testing.T) {
	t.Parallel()
	tests := []struct{
		name string
		SquareMatrix *SquareMatrix
		coords [2]int
		want float64
	}{
		{
			"2x2_1",
			&SquareMatrix{[]float64{1, 2, 3, 4}, 2},
			[2]int{0, 0},
			1,
		},
		{
			"2x2_2",
			&SquareMatrix{[]float64{1, 2, 3, 4}, 2},
			[2]int{1, 0},
			3,
		},
		{
			"3x3",
			&SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
			[2]int{2, 2},
			9,
		},
	}

	for _, test := range tests {
		testname := test.name
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			result := test.SquareMatrix.Get(test.coords[0], test.coords[1])
			if result != test.want {
				t.Fatalf(`Failed to get value at %v. Expected %f, got %f`,
					test.coords, test.want, result)
			}
		})
	}
}

// TestString calles SquareMatrix.String
func TestSqString(t *testing.T) {
	t.Parallel()
	tests := []struct{
		name string
		SquareMatrix *SquareMatrix
		want string
	}{
		{
			"2x2",
			&SquareMatrix{[]float64{1, 2, 3, 4}, 2},
			"[[1.00 2.00]\n[3.00 4.00]]",
		},
		{
			"3x3",
			&SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
			"[[1.00 2.00 3.00]\n[4.00 5.00 6.00]\n[7.00 8.00 9.00]]",
		},
	}

	for _, test := range tests {
		testname := test.name
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			result := test.SquareMatrix.String()
			if result != test.want {
				t.Fatalf("Failed to get string. Expected \n%s, got \n%s",
					test.want, result)
			}
		})
	}
}

// TestAdd calles SquareMatrix.Add 
func TestSqAdd(t *testing.T) {
	t.Parallel()
	A := SquareMatrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		3,
	}

	B := SquareMatrix{
		[]float64{9, 8, 7, 6, 5, 4, 3, 2, 1},
		3,
	}

	C := A.Add(&B)
	expected := SquareMatrix{
		[]float64{10, 10, 10, 10, 10, 10, 10, 10, 10},
		3,
	}

	if !reflect.DeepEqual(C.data, expected.data){
		t.Fatalf(`Expected %s, got %s`, expected.String(), C.String())
	}
}

// TestSet calls SquareMatrix.Set 
func TestSqSet(t *testing.T) {
	t.Parallel()
	SquareMatrix := SquareMatrix{
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		3,
	}

	tests := []struct {
		name string
		coords   [2]int
		value    float64
		expected []float64
	}{
		{"Triangle1", [2]int{0, 0}, 10, []float64{10, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Triangle2", [2]int{1, 1}, 20, []float64{10, 2, 3, 4, 20, 6, 7, 8, 9}},
		{"Triangle3", [2]int{2, 2}, 30, []float64{10, 2, 3, 4, 20, 6, 7, 8, 30}},
	}

	for _, test := range tests {
		testname := test.name
		// Cannot call t.Parallel since we are changing the SquareMatrix
		t.Run(testname, func(t *testing.T) {
			SquareMatrix.Set(test.coords[0], test.coords[1], test.value)
			if !reflect.DeepEqual(SquareMatrix.data, test.expected) {
				t.Fatalf(`Failed to set value at %v. Expected %v, got %v`,
					test.coords, test.expected, SquareMatrix.data)
			}
		})
	}
}

// TestSqSubMatrix calles SquareMatrix.SetSubMatrix
func TestSqSetSubMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		A *SquareMatrix
		subMatrix *SquareMatrix
		coords [2]int
		want *SquareMatrix
	}{
		{
			"2x2",
			&SquareMatrix{[]float64{1, 2, 3, 4}, 2},
			&SquareMatrix{[]float64{5, 6, 7, 8}, 2},
			[2]int{0, 0},
			&SquareMatrix{[]float64{5, 6, 7, 8}, 2},
		},
		{
			"3x3",
			&SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3},
			&SquareMatrix{[]float64{-4, -5, -7, -8}, 2},
			[2]int{1, 0},
			&SquareMatrix{[]float64{1, 2, 3, -4, -5, 6, -7, -8, 9}, 3},
		},
	}

	for _, test := range tests {
		testname := test.name
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			test.A.SetSubMatrix(test.subMatrix, test.coords[0], test.coords[1])
			if !reflect.DeepEqual(test.A.data, test.want.data) {
				t.Fatalf(`Failed to set submatrix at %v. Expected\n%s, got\n%s`,
					test.coords, test.want.String(), test.A.String())
			}
		})
	}
}

// TestMultiply calles SquareMatrix.Multiply 
func TestSqMultiply(t *testing.T) {
	t.Parallel()

	largeTestMatrix := createSquareMatrix(random, 1024)

	tests := []struct{
		name string
		A *SquareMatrix
		B *SquareMatrix
		want *SquareMatrix
	}{
		{
			"2DShouldPass",
			&SquareMatrix{[]float64{1, 2, 3, 4}, 2},
			&SquareMatrix{[]float64{7, 10, 13, 16}, 2},
			&SquareMatrix{[]float64{33, 42, 73, 94}, 2},
		},
		{
			"IdentitySquareMatrix",
			&SquareMatrix{[]float64{1, 0, 0, 0, 1, 0, 0, 0, 1}, 3},
			&SquareMatrix{[]float64{7, 10, 13, 2, 90, 6, 39, 2, 1}, 3},
			&SquareMatrix{[]float64{7, 10, 13, 2, 90, 6, 39, 2, 1}, 3},
		},
		{
			"4x4ShouldPass",
			&SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4},
			&SquareMatrix{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 4},
			&SquareMatrix{[]float64{90, 100, 110, 120, 202, 228, 254, 280, 314, 356, 398, 440, 426, 484, 542, 600}, 4},
		},
		{
			"LargerShouldPass",
			createSquareMatrix(ones, 100),
			createSquareMatrix(identity, 100),
			createSquareMatrix(ones, 100),
		},
		{
			"LargestShouldPass",
			largeTestMatrix,
			createSquareMatrix(identity, 1024),
			largeTestMatrix,
		},
	}

	for _, test := range tests {
		testname := test.name
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			result := test.A.Multiply(test.B)
			if !reflect.DeepEqual(result.data, test.want.data) {
				t.Fatalf("\nTest case FAILED: %s\nExpected\n%s, got\n%s",
					test.name, test.want.String(), result.String())
			}
		})
	}
}

// ===========================================================================
func BenchmarkSqMultiply(b *testing.B) {
	size := 2048

	// Create two size x size matrices with random numbers between 0 and 20
	A := &SquareMatrix{
		data: make([]float64, size*size),
		n: size,
	}
	B := &SquareMatrix{
		data: make([]float64, size*size),
		n: size,
	}

	// Fill the SquareMatrix with random numbers
	for i := 0; i < len(A.data); i++ {
		A.data[i] = rand.Float64() * 20
		B.data[i] = rand.Float64() * 20
	}

    for i := 0; i < b.N; i++ {
        A.Multiply(B)
    }
}