package errors

import (
	"errors"
	"fmt"
)

func ToFormatError(err error, format string, v ...interface{}) error {

	return errors.New(
		fmt.Sprintf(
			"%s\n    %s\n",
			err.Error(),
			fmt.Sprintf(format, v...)))
}

func NewError(format string, v ...interface{}) error {
	return errors.New(
		fmt.Sprintf(format, v...),
	)
}
