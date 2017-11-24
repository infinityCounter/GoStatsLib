package test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/infinityCounter/statlib/src"
)

var (
	seed = rand.NewSource(time.Now().UnixNano())
	gen  = rand.New(seed)
)

func getPositiveIntNonZero(min, max int) int {
	var num int
	for {
		num = gen.Intn(max)
		if num < min {
			continue
		}
		break
	}
	return num
}

func TestNewMatrix(t *testing.T) {
	for count := 0; count < 100; count++ {
		colSize := getPositiveIntNonZero(0, 100)
		rowSize := getPositiveIntNonZero(0, 365)
		shouldC := colSize
		shouldR := rowSize
		if colSize == 0 {
			shouldC = 1
		}
		if rowSize == 0 {
			shouldR = 1
		}
		matrix := src.NewMatrix(colSize, rowSize)
		if matrix.ColCount() != shouldC {
			t.Errorf("Matrix Column size not equal to what was passed, expected: %d, got: %d", shouldC, matrix.ColCount())
		}
		if matrix.RowCount() != shouldR {
			t.Errorf("Matrix Row size not equal to what was passed, expected: %d, got: %d", shouldR, matrix.RowCount())
		}
		if matrix.Data == nil || len(matrix.Data) == 0 {
			t.Errorf("Matrix Data was no allocated")
		} else {
			if len(matrix.Data) != shouldC {
				t.Errorf("Matrix Column size not equal to what was passed, expected: %d, got: %d", shouldC, matrix.ColCount())
			} else {
				if len(matrix.Data[0]) != shouldR {
					t.Errorf("Matrix Row size not equal to what was passed, expected: %d, got: %d", shouldR, matrix.RowCount())
				}
			}
		}
	}
}

func TestAddCol(t *testing.T) {
	colSize := getPositiveIntNonZero(1, 10)
	rowSize := getPositiveIntNonZero(1, 10)
	mtrx := src.NewMatrix(colSize, rowSize)
	addSize := getPositiveIntNonZero(1, 10)
	mtrx.AddCol(addSize)
	shouldBe := colSize + addSize
	if mtrx.ColCount() != shouldBe || len(mtrx.Data) != shouldBe {
		t.Errorf("Matrix Column size not equal to what was passed, expected: %d, got: %d", shouldBe, mtrx.ColCount())
	}
}

func TestAddRow(t *testing.T) {
	colSize := getPositiveIntNonZero(1, 10)
	rowSize := getPositiveIntNonZero(1, 10)
	mtrx := src.NewMatrix(colSize, rowSize)
	addSize := getPositiveIntNonZero(1, 10)
	mtrx.AddRow(addSize)
	shouldBe := rowSize + addSize
	if mtrx.RowCount() != shouldBe {
		t.Errorf("Matrix Row size not equal to what was passed, expected: %d, got: %d", shouldBe, mtrx.RowCount())
	}
	for _, col := range mtrx.Data {
		if len(col) != shouldBe {
			t.Errorf("Matrix Row size not equal to what was passed, expected: %d, got: %d", shouldBe, len(col))
			break
		}
	}
}

func TestSetPosValue(t *testing.T) {
	colSize := 4
	rowSize := 4
	mtrx := src.NewMatrix(colSize, rowSize)
	val := 7.14
	err := mtrx.SetPosValue(0, 2, val)
	if err != nil {
		t.Error(err.Error())
	}
	if v, err := mtrx.GetPosValue(0, 2); err != nil {
		t.Error(err.Error())
	} else if v != val {
		t.Errorf("Matrix value at position not what was expected, expected: %f, actual: %f", val, v)
	}
}

func TestPosValueExpectErr(t *testing.T) {
	colSize := 4
	rowSize := 4
	mtrx := src.NewMatrix(colSize, rowSize)
	err := mtrx.SetPosValue(9, 1, 8)
	if err == nil {
		t.Errorf("Matrix set pos value, no error encountered when setting out of bounds index")
	}
}

func TestDivideAllByFloat(t *testing.T) {
	colSize := 4
	rowSize := 2
	vals := [][]float64{
		{5, 6},
		{9.1, 3.3},
		{1.1, 2.4},
		{9.1, 0},
	}
	mtrx := src.NewMatrix(colSize, rowSize)
	for ccount := 0; ccount < colSize; ccount++ {
		for rcount := 0; rcount < rowSize; rcount++ {
			mtrx.SetPosValue(ccount, rcount, vals[ccount][rcount])
		}
	}
	divisor := 6.0
	mtrx.DivideAllByFloat64(6)
	for cidx, col := range mtrx.Data {
		for ridx, cell := range col {
			if cell != (vals[cidx][ridx] / divisor) {
				t.Errorf("Matrix divide by all failed, expected:%f, actual:%f", (vals[cidx][ridx] / divisor), cell)
				break
			}
		}
	}
}

func TestNormalizeMatrix(t *testing.T) {
	colSize := 4
	rowSize := 2
	vals := [][]float64{
		{5, 6},
		{9.1, 3.3},
		{1.1, 2.4},
		{9.1, 0},
	}
	mtrx := src.NewMatrix(colSize, rowSize)
	for ccount := 0; ccount < colSize; ccount++ {
		for rcount := 0; rcount < rowSize; rcount++ {
			mtrx.SetPosValue(ccount, rcount, vals[ccount][rcount])
		}
	}
	mtrx.NormalizeMatrix()
	for cidx, col := range mtrx.Data {
		for ridx, cell := range col {
			normalized := (vals[cidx][ridx] / vals[cidx][0])
			if cell != normalized {
				t.Errorf("Matrix divide by all failed, expected:%f, actual:%f", normalized, cell)
				break
			}
		}
	}
}

func TestGetCol(t *testing.T) {
	colSize := 4
	rowSize := 2
	vals := [][]float64{
		{5, 6},
		{9.1, 3.3},
		{1.1, 2.4},
		{9.1, 0},
	}
	mtrx := src.NewMatrix(colSize, rowSize)
	for ccount := 0; ccount < colSize; ccount++ {
		for rcount := 0; rcount < rowSize; rcount++ {
			mtrx.SetPosValue(ccount, rcount, vals[ccount][rcount])
		}
	}
	col, err := mtrx.GetCol(0)
	if err != nil {
		t.Errorf("Matrix did not return column")
	}
	if len(col) != len(vals[0]) {
		t.Errorf("Matrix column returned does not match, expected:%v, actual:%v", vals[0], col)
	}
	for idx, val := range vals[0] {
		if val != col[idx] {
			t.Errorf("Matrix column returned does not match, expected:%v, actual:%v", vals[0], col)
			break
		}
	}
}
