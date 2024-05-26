package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/yinebebt/priceestimation/utils"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"time"
)

// Location represents user-specified location data
type Location struct {
	ID        uuid.UUID `json:"id"`
	Country   string    `json:"country"`
	Region    string    `json:"region"`
	Zone      string    `json:"zone"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PriceEstimation represents the price estimation data
type PriceEstimation struct {
	ID          uuid.UUID       `json:"id"`
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	// UserID is unique identifier of the user creating the estimation
	UserID     uuid.UUID `json:"user_id"`
	LocationID uuid.UUID `json:"location_id"`
	Location   Location  `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validations
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Email, validation.Required.Error("email is required"),
			is.EmailFormat.Error("email is not valid")),
		validation.Field(&u.Password, validation.Required.Error("password is required"),
			PassWord.Error("weak password")),
	)
}

var ErrPass = validation.NewError("validation_is_password", "must be strong password")
var PassWord = validation.NewStringRuleWithError(utils.IsPasswordValid, ErrPass)

func (p PriceEstimation) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Price, validation.Required.Error("price is required")),
		validation.Field(&p.LocationID, validation.Required.Error("location is required")),
	)
}

func (l Location) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Country, validation.Required.Error("country is required")),
		validation.Field(&l.Region, validation.Required.Error("Region is required")),
	)
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u LoginRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error("email is required"),
			is.EmailFormat.Error("Invalid email format")),
		validation.Field(&u.Password, validation.Required.Error("password is required")),
	)
}

type LogInResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}
