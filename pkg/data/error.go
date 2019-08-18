package data

import (
	"fmt"
	"strconv"
	"strings"
)

type DataError struct {
	Code    int
	Message string
	Action  string
	Cause   error
}

func (de DataError) Error() string {
	var e string
	format := "%s %s %s %s"

	e = fmt.Sprintf(format, strconv.Itoa(de.Code), de.Message, de.Cause, de.Action)
	e = strings.Replace(e, "%!s(<nil>)", "", -1) // account nils
	e = strings.TrimLeft(e, "0")                 // remove code if it is not set

	for strings.Contains(e, "  ") { // make sure to remove all double spaces if encoutered
		e = strings.Replace(e, "  ", " ", -1) // account for multiple empty strings
	}
	e = strings.Trim(e, " ") // remove trailing/leading whitespace

	return e
}
