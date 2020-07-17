package common

import (
	"errors"
	"fmt"
)

var ErrStringTooLong = errors.New("string too long")
var ErrInvalidLength = errors.New("invalid string length")

type LengthConstraintError [2]int

func (e LengthConstraintError) Error() string {
	return fmt.Sprintf("expected string length between %d and %d", e[0], e[1])
}
