package src

type DataError interface {
	Code() int
	Error() string
}

type dataerr struct {
	code    int
	message string
}

func (de dataerr) Code() int {
	return de.code
}

func (de dataerr) Error() string {
	return de.message
}

var (
	IndexOutOfBoundsError = dataerr{
		code:    0,
		message: "Index specified is out of bounds, cannot acces data structure element at position",
	}

	IndexOrderInvalidError = dataerr{
		code:    1,
		message: "Start index greater than end index, cannot access elements at data structure position",
	}

	DataFrameDuplicateHeadingError = dataerr{
		code:    2,
		message: "DataFrame cannot contain duplicate column/row heading",
	}
)
