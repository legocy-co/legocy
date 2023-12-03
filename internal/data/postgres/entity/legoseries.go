package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSeriesPostgres struct {
	Model
	Name string `gorm:"unique"`
}

func (lsp LegoSeriesPostgres) TableName() string {
	return "lego_series"
}

func FromLegoSeriesValueObject(s *models.LegoSeriesValueObject) *LegoSeriesPostgres {
	return &LegoSeriesPostgres{
		Name: s.Name,
	}
}

func FromLegoSeries(s *models.LegoSeries) *LegoSeriesPostgres {
	return &LegoSeriesPostgres{
		Name: s.Name,
	}
}

func (s *LegoSeriesPostgres) ToLegoSeries() *models.LegoSeries {
	return &models.LegoSeries{
		ID:   int(s.ID),
		Name: s.Name,
	}
}
