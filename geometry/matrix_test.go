package geometry

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func CreateMatrixWithTestData() *Matrix {
	a := make([]float64, 16)

	a[0] = 2.0
	a[1] = 2.3
	a[2] = 5.3
	a[3] = 9.2

	a[4] = 7.2
	a[5] = 3.5
	a[6] = 1.6
	a[7] = 3.6

	a[8] = 5.4
	a[9] = 4.2
	a[10] = 9.4
	a[11] = 8.1

	a[12] = 4.4
	a[13] = 5.6
	a[14] = 7.1
	a[15] = 3.6

	m := CreateMatrixWithElements(4, 4, a)

	/* Result if you multiply the matrix by itself
	* 89.660  86.430 129.420 102.730
	* 64.080  55.690  84.360 104.760
	* 127.440 111.960 181.210 170.100
	* 103.300  79.700 124.580 131.110
	 */

	return m
}

func TestCreateMatrix(t *testing.T) {
	m := CreateMatrix(4, 4)

	if m.rows != 4 || m.cols != 4 {
		t.Fatal("Wrong number of rows and columns in the matrix")
	}

	for m, j := range m.elements {
		if j != 0 {
			t.Fatalf("Bad initialization of the matrix: %d", m)
		}
	}
}

func TestCreateMatrixWithElements(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := make([]float64, 16)

	for i, _ := range arr {
		arr[i] = rand.Float64()
	}

	m := CreateMatrixWithElements(4, 4, arr)

	if len(arr) != len(m.elements) {
		t.Fatal("Matrix size doesn't match size")
	}

	for i := 0; i < 16; i++ {
		if arr[i] != m.elements[i] {
			t.Fatal("Matrix elements do not match: ", arr[i], m.elements[i])
		}
	}

	for i, _ := range arr {
		arr[i] = -1.0
	}

	for i := 0; i < 16; i++ {
		if arr[i] == m.elements[i] {
			t.Fatal("Matrix elements do not match: ", arr[i], m.elements[i])
		}
	}
}

func TestMatrixMultiplication(t *testing.T) {
	m := CreateMatrixWithTestData()
	n := CreateMatrixWithTestData()

	j, _ := Multiply(m, n)

	if fmt.Sprintf("%.2f", j.GetElementAt(0, 0)) != "89.66" ||
		fmt.Sprintf("%.2f", j.GetElementAt(0, 1)) != "86.43" ||
		fmt.Sprintf("%.2f", j.GetElementAt(0, 2)) != "129.42" ||
		fmt.Sprintf("%.2f", j.GetElementAt(0, 3)) != "102.73" {
		t.Fatal("Could not multiply matrices correctly: ", m.ToString())
	}
}

func TestIdentityMatrix(t *testing.T) {
	m := CreateIdentityMatrix(3, 3)
	if m.GetElementAt(0, 0) != 1 || m.GetElementAt(1, 1) != 1 || m.GetElementAt(2, 2) != 1 || m.GetElementAt(0, 1) != 0 {
		t.Fatal("Identity matrix could not be created: ", m.ToString())
	}
}
