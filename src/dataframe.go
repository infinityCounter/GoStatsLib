package src

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type DataFrameLookup struct {
	Col, Row interface{}
}

func (dfl DataFrameLookup) GetColLookup() fmt.Stringer {
	return NewStringer(dfl.Col)
}

func (dfl DataFrameLookup) GetColLookupString() string {
	return NewStringer(dfl.Col).String()
}

func (dfl DataFrameLookup) GetRowLookup() fmt.Stringer {
	return NewStringer(dfl.Row)
}

func (dfl DataFrameLookup) GetRowLookupString() string {
	return NewStringer(dfl.Row).String()
}

type DFL DataFrameLookup

type DataFrame struct {
	RowLabels, ColHeadings []fmt.Stringer
	*Matrix
}

func NewDataFrame(dfc DataFrameConfiguration) (*DataFrame, DataError) {
	df := &DataFrame{}
	if dfc.FromCSV() {
		headingsFromCSV := false
		headingOption := dfc.GetHeadingsOption()
		if headingOption == useFirstAsHeadings {
			headingsFromCSV = true
		}
		df.loadFromCSV(dfc.GetCSVPath(), headingsFromCSV)

		if headingOption == useNoHeadings {
			//If no headings provided just make them regular ints
			for count := 0; count < df.colCount; count++ {
				df.ColHeadings = append(df.ColHeadings, NewStringer(count))
			}

			for count := 0; count < df.rowCount; count++ {
				df.RowLabels = append(df.RowLabels, NewStringer(count))
			}
		}

	} else {
		if dfc.GetHeadingsOption() == useHeadingArgs {
			df.ColHeadings = dfc.GetColHeadings()
			df.RowLabels = dfc.GetRowLabels()
		} else {
			df.ColHeadings = []fmt.Stringer{}
			df.RowLabels = []fmt.Stringer{}
		}
		df.Matrix = NewMatrix(len(df.ColHeadings), len(df.RowLabels))
	}
	if !df.ValidateColHeaders() {
		return nil, DataFrameDuplicateHeading
	}
	return df, nil
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
	for idx, heading := range df.RowLabels {
		for idx2, h2 := range df.RowLabels {
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
			if len(df.RowLabels) == 0 {
				for _, heading := range entry {
					df.ColHeadings = append(df.ColHeadings, NewStringer(heading))
				}
			}
			df.RowLabels = append(df.RowLabels, NewStringer(entry[0]))
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

func (df DataFrame) GetColHeadingIndex(needle fmt.Stringer) int {
	for idx, heading := range df.ColHeadings {
		if heading.String() == needle.String() {
			return idx
		}
	}
	return -1
}

func (df DataFrame) GetRowLabelIndex(needle fmt.Stringer) int {
	for idx, heading := range df.RowLabels {
		if heading.String() == needle.String() {
			return idx
		}
	}
	return -1
}

func (df DataFrame) GetValue(dfl DataFrameLookup) (float64, DataError) {
	if dfl.Col == nil || dfl.Row == nil {
		return 0, DataFrameInsufficientLookupArguments
	}
	cIdx := df.GetColHeadingIndex(dfl.GetColLookup())
	rIdx := df.GetRowLabelIndex(dfl.GetRowLookup())
	return df.Matrix.GetPosValue(cIdx, rIdx)
}

func (df DataFrame) SetValue(dfl DataFrameLookup, val float64) DataError {
	if dfl.Col == nil || dfl.Row == nil {
		return DataFrameInsufficientLookupArguments
	}
	cIdx := df.GetColHeadingIndex(dfl.GetColLookup())
	rIdx := df.GetRowLabelIndex(dfl.GetRowLookup())
	return df.Matrix.SetPosValue(cIdx, rIdx, val)
}
