package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

// Location represents user-specified location data
type Location struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	Woreda  string `json:"woreda"`
	City    string `json:"city"`
}

// PriceEstimation represents the price estimation data
type PriceEstimation struct {
	ID          uuid.UUID       `json:"id"`
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	// UserID is unique identifier of the user creating the estimation
	UserID    uuid.UUID `json:"user_id"`
	Location  Location  `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
