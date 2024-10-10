package entity

import (
	"errors"
	"time"

	"github.com/Kayky18/API_FULLCYCLE/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id-is-required")
	ErrIdIsInvalid     = errors.New("id-is-invalid")
	ErrNameIsRequired  = errors.New("name-is-required")
	ErrPriceIsRequired = errors.New("price-is-required")
	ErrPriceInvalid    = errors.New("invalid-price")
)

type Product struct {
	ID        entity.ID `json:"id"`         // id
	Name      string    `json:"name"`       // name
	Price     float64   `json:"price"`      // price
	CreatedAt time.Time `json:"created_at"` // created_at
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIdIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrPriceInvalid
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {

	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}
	return product, nil

}
