package service

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	"github.com/legocy-co/legocy/internal/domain/calculator/repository"
)

type LegoSetValuationService struct {
	r repository.LegoSetValuationRepository
}

func NewLegoSetValuationService(r repository.LegoSetValuationRepository) LegoSetValuationService {
	return LegoSetValuationService{r: r}
}

func (s LegoSetValuationService) GetLegoSetValuations(c context.Context, legoSetId int) ([]models.LegoSetValuation, *errors.AppError) {
	return s.r.GetLegoSetValuationsList(c, legoSetId)
}
