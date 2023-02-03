package marketplace

import (
	"context"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemRepository interface {
	GetMarketItems(c context.Context) ([]*models.MarketItem, error)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, error)
	GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, error)
	GetMarketItemsBySeller(c context.Context, sellerID int) ([]*models.MarketItem, error)
	CreateMarketItem(c context.Context, item *models.MarketItemBasic) error
	DeleteMarketItem(c context.Context, id int) error
}
