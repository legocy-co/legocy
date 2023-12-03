package collections

import (
	"github.com/legocy-co/legocy/internal/app/errors"
)

var (
	ErrValuationNotFound = errors.NewAppError(errors.NotFoundError, "LegoSet Valuation Not Found")
)
