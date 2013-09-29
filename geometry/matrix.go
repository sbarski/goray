package geometry

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Matrix struct {
	//number of rows, cols and the length of the row (step)
	rows, cols int

	//elements [i*rows+j] => row i, col j
	elements []float64
}

func CreateMatrix(rows int, cols int) *Matrix {
	m := new(Matrix)
	m.rows = rows
	m.cols = cols
	m.elements = make([]float64, rows*cols)

	return m
}

func CreateIdentityMatrix(rows int, cols int) *Matrix {
	m := CreateMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j {
				m.SetElementAt(i, j, 1)
			}
		}
	}

	return m
}

func CreateRandomMatrix(rows int, cols int) *Matrix {
	rand.Seed(time.Now().UTC().UnixNano())
	m := CreateMatrix(rows, cols)

	for i, _ := range m.elements {
		m.elements[i] = rand.Float64()
	}

	return m
}

func CreateMatrixWithElements(rows int, cols int, elements []float64) *Matrix {
	m := CreateMatrix(rows, cols)

	copy(m.elements, elements)

	return m
}

func Multiply(a *Matrix, m *Matrix) (*Matrix, error) {
	if m.rows != a.rows {
		return nil, errors.New("Matrix: row size must be the same")
	}

	if m.cols != a.cols {
		return nil, errors.New("Matrix: column size must be the same")
	}

	n := CreateMatrix(m.rows, m.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			v := a.GetElementAt(i, 0)*m.GetElementAt(0, j) +
				a.GetElementAt(i, 1)*m.GetElementAt(1, j) +
				a.GetElementAt(i, 2)*m.GetElementAt(2, j) +
				a.GetElementAt(i, 3)*m.GetElementAt(3, j)

			n.SetElementAt(i, j, v)
		}
	}

	return n, nil
}

func (m *Matrix) SetElementAt(i int, j int, value float64) {
	m.elements[i*m.rows+j] = value
}

func (m *Matrix) GetElementAt(i int, j int) float64 {
	return m.elements[i*m.rows+j]
}

func (m *Matrix) IsEqualTo(a *Matrix) bool {
	a_slice := fmt.Sprintf("%v", a.elements)
	m_slice := fmt.Sprintf("%v", m.elements)

	return a_slice == m_slice
}

func (m *Matrix) ToString() string {
	return fmt.Sprintf("%v", m.elements)
}

func (m *Matrix) Matrix2D() [][]float64 {
	a := make([][]float64, m.rows)

	for i := 0; i < m.rows; i++ {
		a[i] = m.elements[i*m.rows : i*m.rows+m.cols]
	}

	return a
}

func (m *Matrix) Matrix1D() []float64 {
	a := make([]float64, m.rows*m.cols)

	a = m.elements[0 : m.rows*m.cols]

	return a
}
