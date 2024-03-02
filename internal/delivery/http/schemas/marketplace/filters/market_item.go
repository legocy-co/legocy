package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/app/errors"
	legoFilters "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego/filters"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
)

func GetMarketItemFilterCritera(ctx *gin.Context) (*domain.MarketItemFilterCriteria, *errors.AppError) {
	var filterDTO MarketItemFilterDTO

	if err := ctx.BindQuery(&filterDTO); err != nil {
		return nil, nil
	}

	return filterDTO.ToCriteria()
}

type MarketItemFilterDTO struct {
	PriceGTE  *float64                      `form:"price_gte" json:"price_gte"`
	PriceLTE  *float64                      `form:"price_lte" json:"price_lte"`
	SetStates []string                      `form:"set_state__in" json:"set_state__in"`
	Locations []string                      `form:"location__in" json:"location__in"`
	LegoSet   *legoFilters.LegoSetFilterDTO `form:"lego_set" json:"lego_set"`
}

func (dto *MarketItemFilterDTO) ToCriteria() (*domain.MarketItemFilterCriteria, *errors.AppError) {

	if dto.LegoSet == nil {
		return domain.NewMarketItemFilterCriteria(
			dto.PriceGTE,
			dto.PriceLTE,
			dto.SetStates,
			dto.Locations,
			nil,
		)
	}

	return domain.NewMarketItemFilterCriteria(
		dto.PriceGTE,
		dto.PriceLTE,
		dto.SetStates,
		dto.Locations,
		dto.LegoSet.ToCriteria(),
	)
}