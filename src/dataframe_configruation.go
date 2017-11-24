package src

import "fmt"

type headingsOption int

const (
	useNoHeadings headingsOption = iota
	useHeadingArgs
	useFirstAsHeadings
)

type DataFrameConfiguration struct {
	headingStyle           headingsOption
	fromCSV                bool
	filePath               string
	rowLabels, colHeadings []fmt.Stringer
}

func NewDFConfig() DataFrameConfiguration {
	return DataFrameConfiguration{
		headingStyle: useNoHeadings,
		fromCSV:      false,
		rowLabels:    []fmt.Stringer{},
		colHeadings:  []fmt.Stringer{},
	}
}

func (dfc DataFrameConfiguration) GetHeadingsOption() headingsOption {
	return dfc.headingStyle
}

func (dfc *DataFrameConfiguration) UseNoHeadings() {
	dfc.headingStyle = useNoHeadings
}

func (dfc *DataFrameConfiguration) UseHeadingArgs() {
	dfc.headingStyle = useHeadingArgs
}

func (dfc *DataFrameConfiguration) UseFirstAsHeadings() {
	dfc.headingStyle = useFirstAsHeadings
}

func (dfc *DataFrameConfiguration) GetCSVPath() string {
	return dfc.filePath
}

func (dfc *DataFrameConfiguration) SetCSVPath(path string) {
	dfc.filePath = path
	dfc.fromCSV = true
}

func (dfc *DataFrameConfiguration) FromCSV() bool {
	return dfc.fromCSV
}

func (dfc *DataFrameConfiguration) SetRowLabels(headings []fmt.Stringer) {
	dfc.rowLabels = headings
	if dfc.headingStyle != useFirstAsHeadings {
		dfc.UseHeadingArgs()
	}
}

func (dfc *DataFrameConfiguration) SetColHeadings(headings []fmt.Stringer) {
	dfc.colHeadings = headings
	if dfc.headingStyle != useFirstAsHeadings {
		dfc.UseHeadingArgs()
	}
}

func (dfc *DataFrameConfiguration) SetHeadings(col, row []fmt.Stringer) {
	dfc.SetColHeadings(col)
	dfc.SetRowLabels(row)
}

func (dfc DataFrameConfiguration) GetColHeadings() []fmt.Stringer {
	return dfc.colHeadings
}

func (dfc DataFrameConfiguration) GetRowLabels() []fmt.Stringer {
	return dfc.rowLabels
}
