package cart

import (
	"github.com/pallat/ddd/product"
)

type Cart struct {
	Items []*product.Item
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(item *product.Item) {
	c.Items = append(c.Items, item)
}
