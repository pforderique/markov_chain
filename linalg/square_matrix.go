package linalg

import (
	"fmt"
	"strings"
	"sync"
)

// SquareMatrix is a struct that represents an n x n matrix.
type SquareMatrix struct {
	data []float64
	n int
}


func (A SquareMatrix) Size() int {
	return A.n*A.n
}

func (A SquareMatrix) isValidIndex(i, j int) bool {
	return i < A.n && j < A.n && i >= 0 && j >= 0
}

func (A SquareMatrix) Get(i, j int) float64 {
	if !A.isValidIndex(i, j) {
		panic(fmt.Sprintf(
			"Index out of range for SquareMatrix of size %d", A.n))
	}

    return A.data[i*A.n + j]
}

func (A *SquareMatrix) Set(x,y int, value float64) {
	if !A.isValidIndex(x, y) {
		panic(fmt.Sprintf(
			"Index out of range for SquareMatrix of size %d", A.n))
	}
	A.data[x*A.n + y] = value
}

func (A SquareMatrix) String() string {
	strPieces := []string{}
	for i:=0; i<A.n; i++ {
		subStringPieces := []string{}
		for j:=0; j<A.n; j++ {
			subStringPieces = append(
				subStringPieces, fmt.Sprintf("%.2f", A.Get(i, j)))
		}
		strPieces = append(
			strPieces, fmt.Sprintf("[%s]", strings.Join(subStringPieces, " ")))
	}
	return fmt.Sprintf("[%s]", strings.Join(strPieces, "\n"))
}

func (A *SquareMatrix) Add(B *SquareMatrix) *SquareMatrix {
	if A.n != B.n {
		panic(fmt.Sprintf(
			"SquareMatrix dimensions (%dx%d)+(%dx%d) cannot be added", 
			A.n, A.n, B.n, B.n))
	}

	C := SquareMatrix{make([]float64, A.Size()), A.n}
	for i:=0; i<A.n; i++ {
		for j:=0; j<A.n; j++ {
			C.Set(i, j, A.Get(i, j) + B.Get(i, j))
		}
	}
	return &C
}

func (A *SquareMatrix) Multiply(B *SquareMatrix) *SquareMatrix {
	if A.n != B.n {
		panic(fmt.Sprintf(
			"SquareMatrix dimensions (%dx%d)*(%dx%d) cannot be multiplied",
			A.n, A.n, B.n, B.n))
	}
	return SquareMatrixMultiplyDense(A, B)
}

func SquareMatrixMultiplySimple(A *SquareMatrix, B *SquareMatrix) *SquareMatrix {
	C := SquareMatrix{make([]float64, A.Size()), A.n}
	for i:=0; i<A.n; i++ {
		for k:=0; k<A.n; k++ {
			var value float64 = 0
			for j:=0; j<A.n; j++ {
				value += A.Get(i, j) * B.Get(j, k)
			}
			C.Set(i, k, value)
		}
	}
	return &C
}

// SetSubMatrix sets the values of subMatrix starting at (i, j) in A
func (A *SquareMatrix) SetSubMatrix(subMatrix *SquareMatrix, i, j int) {
	for a:=i; a<i+subMatrix.n; a++ {
		for b:=j; b<j+subMatrix.n; b++ {
			A.Set(a, b, subMatrix.Get(a-i, b-j))
		}
	}
}

// getSubMatrix returns a the s x s subMatrix starting at (i, j) in A
func getSubMatrix(A *SquareMatrix, i, j, s int) *SquareMatrix {
	subMatrix := SquareMatrix{make([]float64, s*s), s}
	for a:=0; a<s; a++ {
		for b:=0; b<s; b++ {
			subMatrix.Set(a, b, A.Get(i+a, j+b))
		}
	}
	return &subMatrix
}

// chooseP returns the optimal value of p for a given n
func chooseP(n int) int {
	// TODO: Improve on this
	if n % 2 == 1 || n < 100 {
		return 1
	}

	if n == 100 || n == 500 {
		return 4
	} else if n == 1024 {
		return 16
	} else if n == 2048 {
		return 32
	}

	p := 1
	for n >= 100 {
		n /= 10
		p *= 2
	}
	return p
}

func SquareMatrixMultiplyDense(A *SquareMatrix, B *SquareMatrix) *SquareMatrix {
	// Split A and B into p*p submatrices each of size n/p x n/p
	n := A.n
	p := chooseP(n)

	// fmt.Printf("n: %d, p: %d\n", n, p)

	if p == 1 {
		return SquareMatrixMultiplySimple(A, B)
	}

	// Create a new result SquareMatrix C of size n x n
	C := SquareMatrix{make([]float64, A.Size()), n}

	// Store submatrices of A and B in maps
	subMatricesA := make(map[string]*SquareMatrix)
	AMutex := sync.RWMutex{}
	subMatricesB := make(map[string]*SquareMatrix)
	BMutex := sync.RWMutex{}
	subMatrixDones := make(chan bool)

	// Create goroutines to calculate submatrices concurrently
	for i:=0; i<p; i++ {
		for j:=0; j<p; j++ {
			go func(i, j, p int) {
				AMutex.Lock()
				Aij := getSubMatrix(A, i*n/p, j*n/p, n/p)
				subMatricesA[fmt.Sprintf("%d,%d", i, j)] = Aij
				AMutex.Unlock()
				subMatrixDones <- true
			}(i, j, p)
			
			go func(i, j, p int) {
				BMutex.Lock()
				Bij := getSubMatrix(B, i*n/p, j*n/p, n/p)
				subMatricesB[fmt.Sprintf("%d,%d", i, j)] = Bij
				BMutex.Unlock()
				subMatrixDones <- true
			}(i, j, p)
		}
	}

	// Wait for all of A and B's p^2 submatrices to be calculated
	for i:=0; i<2*p*p; i++ {
		<-subMatrixDones
	}

	// Print submatrices
	// fmt.Println("Submatrices A:")
	// count := 0
	// for key, value := range subMatricesA {
	// 	if count == 10 {
	// 		break
	// 	}
	// 	fmt.Printf("%s:\n%s\n    ---    \n", key, value.String())
	// 	count++
	// }
	// fmt.Println("Submatrices B:")
	// for key, value := range subMatricesB {
	// 	fmt.Printf("%s:\n%s\n    ---    \n", key, value.String())
	// }

	// Create a channel to store results of submatrix multiplications
	subMatrixResults := make(chan struct{i, j int; subMatrix *SquareMatrix})

	// Create goroutines to multiply submatrices concurrently
	for i:=0; i<p; i++ {
		for j:=0; j<p; j++ {
			go func(i, j, p int) {
				Cij := &SquareMatrix{make([]float64, (n/p)*(n/p)), n/p}
				for k:=0; k<p; k++ {
					// TODO: Do I need to use locks here?
					Aik := subMatricesA[fmt.Sprintf("%d,%d", i, k)]
					Bkj := subMatricesB[fmt.Sprintf("%d,%d", k, j)]
					Cij = SquareMatrixMultiplySimple(Aik, Bkj).Add(Cij)
				}
				subMatrixResults <- struct{i, j int; subMatrix *SquareMatrix}{i, j, Cij}
			}(i, j, p)
		}
	}

	// Process all p^2 submatrix results as they come in
	for i:=0; i<p*p; i++ {
		result := <-subMatrixResults
		C.SetSubMatrix(result.subMatrix, result.i*(n/p), result.j*(n/p))
	}

	return &C
}
