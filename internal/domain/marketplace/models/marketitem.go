package marketplace

import (
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	users "github.com/legocy-co/legocy/internal/domain/users/models"
)

type MarketItem struct {
	ID          int
	LegoSet     lego.LegoSet
	Seller      users.User
	Price       float32
	Location    string
	SetState    string // lego.SetStateBrandNew / etc.
	Status      string // lego.ListingStatusCheckRequired / lego.ListingStatusActive / etc.
	Description string
	Images      []*MarketItemImage
}

type MarketItemValueObject struct {
	LegoSetID   int
	SellerID    int
	Price       float32
	CurrencyID  int
	Location    string
	Status      string
	SetState    string // SetStateBrandNew / etc.
	Description string
}
