package repository

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetRepository interface {
	CreateLegoSet(c context.Context, s *models.LegoSetValueObject) *errors.AppError
	GetLegoSets(c context.Context) ([]*models.LegoSet, *errors.AppError)
	GetLegoSetByID(c context.Context, id int) (*models.LegoSet, *errors.AppError)
	DeleteLegoSet(c context.Context, id int) *errors.AppError
}
