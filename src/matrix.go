package src

type Matrix struct {
	rowCount, colCount int
	// Data is indexed [Column][Row]
	Data [][]float64
}

func NewMatrix(numCols, numRows int) *Matrix {
	if numCols == 0 {
		numCols = 1
	}
	if numRows == 0 {
		numRows = 1
	}
	data := make([][]float64, numCols)
	for idx, _ := range data {
		// The value at the column index is the first val of the row in the matrix
		data[idx] = make([]float64, numRows)
	}
	return &Matrix{
		rowCount: numRows,
		colCount: numCols,
		Data:     data,
	}
}

func (mtrx Matrix) ColCount() int {
	return mtrx.colCount
}

func (mtrx Matrix) RowCount() int {
	return mtrx.rowCount
}

// AddRow a appends the number of rows dictated by the passed argument to the Matrix
func (mtrx *Matrix) AddRow(numRows int) {
	if numRows < 0 {
		panic("statlib.Matrix.AddRow, cannot add a negative number of rows")
	} else if numRows == 0 {
		return
	}
	for idx, col := range mtrx.Data {
		mtrx.Data[idx] = append(col, make([]float64, numRows)...)
	}
	mtrx.rowCount = mtrx.rowCount + numRows
}

func (mtrx *Matrix) AddCol(numCols int) {
	if numCols < 0 {
		panic("statlib.Matrix.AddCol, cannot add a negative number of columns")
	} else if numCols == 0 {
		return
	}
	for count := int(0); count < numCols; count++ {
		// the columns are just the first index of each row
		rows := make([]float64, mtrx.rowCount)
		mtrx.Data = append(mtrx.Data, rows)
	}
	mtrx.colCount = mtrx.colCount + numCols
}

func (mtrx Matrix) GetCol(pos int) ([]float64, DataError) {
	if pos < 0 || pos >= mtrx.colCount {
		return []float64{}, IndexOutOfBounds
	}
	return mtrx.Data[pos], nil
}

func (mtrx Matrix) GetRow(pos int) ([]float64, DataError) {
	if pos < 0 || pos >= mtrx.rowCount {
		return []float64{}, IndexOutOfBounds
	}
	var row []float64
	for _, col := range mtrx.Data {
		row = append(row, col[pos])
	}
	return row, nil
}

func (mtrx Matrix) GetPosValue(col, row int) (float64, DataError) {
	if (col < 0 || row < 0) || (col >= mtrx.colCount || row >= mtrx.rowCount) {
		return 0, IndexOutOfBounds
	}
	return mtrx.Data[col][row], nil
}

func (mtrx *Matrix) SetRow(rowIndex int, vals []float64) DataError {
	if rowIndex < 0 || rowIndex >= mtrx.rowCount {
		return IndexOutOfBounds
	}
	lenDiff := len(vals) - mtrx.ColCount()
	if lenDiff < 0 {
		abs := -lenDiff
		vals = append(vals, make([]float64, abs)...)
	} else if lenDiff > 0 {
		vals = vals[:mtrx.ColCount()]
	}
	for idx, col := range mtrx.Data {
		col[rowIndex] = vals[idx]
	}
	return nil
}

func (mtrx *Matrix) SetPosValue(col, row int, value float64) DataError {
	if (col < 0 || row < 0) || (col >= mtrx.colCount || row >= mtrx.rowCount) {
		return IndexOutOfBounds
	}
	mtrx.Data[col][row] = value
	return nil
}

// GetSubMatrix returns a Matrix containing the data from the subrange specified, columns and rows are inclusive
func (mtrx Matrix) GetSubMatrix(fromCol, toCol, fromRow, toRow int) (*Matrix, DataError) {
	if (fromCol < 0 || toCol < 0 || fromRow < 0 || toRow < 0) ||
		(fromCol >= mtrx.colCount || toCol >= mtrx.colCount || fromRow >= mtrx.rowCount || toRow >= mtrx.rowCount) {
		return nil, IndexOutOfBounds
	}
	if fromCol > toCol || fromRow > toRow {
		return nil, IndexOrderInvalid
	}
	data := mtrx.Data[fromCol : toCol+1]
	for _, col := range data {
		col = col[fromRow : toRow+1]
	}
	sub := NewMatrix(fromCol-toCol+1, fromRow-toRow+1)
	sub.Data = data
	return sub, nil
}

func (mtrx *Matrix) DivideAllByFloat64(val float64) {
	for cidx, col := range mtrx.Data {
		for ridx, cell := range col {
			mtrx.Data[cidx][ridx] = cell / val
		}
	}
}

func (mtrx *Matrix) NormalizeMatrix() {
	for cIdx, col := range mtrx.Data {
		firstVal := col[0]
		for rIdx, cell := range col {
			mtrx.Data[cIdx][rIdx] = cell / firstVal
		}
	}
}
