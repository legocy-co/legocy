package calculator

import (
	"github.com/legocy-co/legocy/internal/app/errors"
)

var (
	ErrLegoSetValuationNotFound = errors.NewAppError(errors.NotFoundError, "Lego Set Valuation Not Found")
)
