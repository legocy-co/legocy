package postgres

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type MarketItemPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemPostgresRepository(conn d.DataBaseConnection) MarketItemPostgresRepository {
	return MarketItemPostgresRepository{conn: conn}
}

func (r MarketItemPostgresRepository) GetMarketItems(
	c context.Context) ([]*models.MarketItem, *errors.AppError) {

	var itemsDB []*entities.MarketItemPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	res := db.Model(&entities.MarketItemPostgres{}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").Preload("Images").
		Find(&itemsDB, "status = 'ACTIVE'")
	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &appErr
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemsAuthorized(
	c context.Context, userID int) ([]*models.MarketItem, *errors.AppError) {

	var itemsDB []*entities.MarketItemPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	res := db.Model(&entities.MarketItemPostgres{}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").Preload("Images").
		Find(&itemsDB, "user_postgres_id <> ? and status = 'ACTIVE'", userID)
	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &appErr
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemByID(
	c context.Context, id int) (*models.MarketItem, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	query := db.Preload("Seller").Preload("Seller.Images").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").Preload("Images").
		Find(&entity, "id = ? and status = 'ACTIVE'", id)

	if query.RowsAffected == 0 {
		return nil, &e.ErrMarketItemsNotFound
	}

	return entity.ToMarketItem(), nil
}

func (r MarketItemPostgresRepository) GetMarketItemsBySellerID(
	c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError) {

	var itemsDB []*entities.MarketItemPostgres
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	result := db.Model(&entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").Preload("Images").
		Find(&itemsDB, "user_postgres_id = ? and status = 'ACTIVE'", sellerID)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return nil, &appErr
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemSellerID(
	c context.Context, id int) (int, *errors.AppError) {

	var count int

	db := r.conn.GetDB()
	if db == nil {
		return count, &d.ErrConnectionLost
	}

	err := db.Model(entities.MarketItemPostgres{}).Where(
		"id=?", id).Select("user_postgres_id").First(&count).Error
	if err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return count, &appErr
	}

	return count, nil
}

func (r MarketItemPostgresRepository) GetSellerMarketItemsAmount(
	c context.Context, sellerID int) (int64, *errors.AppError) {

	var count int64

	db := r.conn.GetDB()
	if db == nil {
		return count, &d.ErrConnectionLost
	}

	res := db.Model(
		entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).Count(&count)

	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return count, &appErr
	}

	return count, nil
}

func (r MarketItemPostgresRepository) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemValueObject(item)
	if entity == nil {
		return &d.ErrItemNotFound
	}

	result := tx.Create(&entity)

	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		tx.Rollback()
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r MarketItemPostgresRepository) DeleteMarketItem(c context.Context, id int) *errors.AppError {

	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	result := db.Delete(entities.MarketItemPostgres{}, id)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}

func (r MarketItemPostgresRepository) UpdateMarketItemByID(
	c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, id)
	if entity == nil {
		return nil, &e.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItem(*item)
	db.Save(entityUpdated)

	return r.GetMarketItemByID(c, id)
}
