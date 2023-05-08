package schema

import (
	"errors"
	"strings"
)

func UpperCaseValidator(value string) error {
	if strings.ToUpper(value) != value {
		return errors.New("should be upper case")
	}
	return nil
}
