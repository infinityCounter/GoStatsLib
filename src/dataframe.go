package src

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type DataFrame struct {
	RowHeadings, ColHeadings []fmt.Stringer
	*Matrix
}

func NewDataFrame(dfc DataFrameConfiguration) (*DataFrame, DataError) {
	df := &DataFrame{}
	if dfc.FromCSV() {
		headingsFromCSV := false
		if dfc.GetHeadingsOption() == useFirstAsHeadings {
			headingsFromCSV = true
		}
		df.loadFromCSV(dfc.GetCSVPath(), headingsFromCSV)

	} else {
		if dfc.GetHeadingsOption() == useHeadingArgs {
			df.ColHeadings = dfc.GetColHeadings()
			df.RowHeadings = dfc.GetRowHeadings()
		} else {
			df.ColHeadings = []fmt.Stringer{}
			df.RowHeadings = []fmt.Stringer{}
		}
		df.Matrix = NewMatrix(len(df.ColHeadings), len(df.RowHeadings))
	}
	if !df.ValidateColHeaders() {
		return nil, DataFrameDuplicateHeadingError
	}
	df, nil
}

func (df DataFrame) ValidateColHeaders() bool {
	for idx, heading := range df.ColHeadings {
		for idx2, h2 := range df.ColHeadings {
			if idx == idx2 {
				continue
			}
			if heading.String() == h2.String() {
				return false
			}
		}
	}
	return true
}

func (df DataFrame) ValidateRowHeaders() bool {
	for idx, heading := range df.RowHeadings {
		for idx2, h2 := range df.RowHeadings {
			if idx == idx2 {
				continue
			}
			if heading.String() == h2.String() {
				return false
			}
		}
	}
	return true
}

func (df DataFrame) ValidateHeadings() bool {
	return df.ValidateColHeaders() && df.ValidateRowHeaders()
}

func (df *DataFrame) loadFromCSV(filePath string, useFirstAsHeadings bool) {
	df.Matrix = nil
	file, err := os.Open(filePath)
	check(err)
	reader := csv.NewReader(file)
	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if useFirstAsHeadings {
			if len(df.RowHeadings) == 0 {
				for _, heading := range entry {
					df.ColHeadings = append(df.ColHeadings, NewStringer(heading))
				}
			}
			df.RowHeadings = append(df.RowHeadings, NewStringer(entry[0]))
		}
		lenElements := len(entry)
		if df.Matrix == nil {
			df.Matrix = NewMatrix(lenElements, 1)
			continue
		} else {
			df.Matrix.AddRow(1)
		}
		floats := make([]float64, lenElements)
		for count := 0; count < lenElements; count++ {
			floats[count] = StringToFloat64(entry[count])
		}
		df.SetRow(df.RowCount()-1, floats)
	}
	// end file read
	file.Close()
}
