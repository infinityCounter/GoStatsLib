package src

import (
	"fmt"
	"strconv"
)

type Stringer struct {
	Val interface{}
}

func NewStringer(val interface{}) Stringer {
	return Stringer{
		Val: val,
	}
}

func (str Stringer) String() string {
	return fmt.Sprintf("%v", str.Val)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringToFloat64(str string) float64 {
	fl, _ := strconv.ParseFloat(str, 64)
	return fl
}
