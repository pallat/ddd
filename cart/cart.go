package cart

import (
	"github.com/pallat/ddd/product"
)

type Cart struct {
	Items map[*product.Item]int64
}

func NewCart() *Cart {
	return &Cart{Items: map[*product.Item]int64{}}
}

func (c *Cart) AddItem(item *product.Item, quantity int64) {
	c.Items[item] = quantity
}
