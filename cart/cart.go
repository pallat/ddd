package cart

import (
	"github.com/pallat/ddd/product"
)

type Cart struct {
	Items map[product.Item]int64
}

func NewCart() *Cart {
	return &Cart{Items: map[product.Item]int64{}}
}

func (c *Cart) AddItem(item *product.Item, quantity int64) {
	c.Items[*item] = quantity
}

func (c *Cart) RemoveItem(item *product.Item, quantity int64) {
	if c.Items[*item] < quantity {
		quantity = c.Items[*item]
	}

	c.Items[*item] -= quantity

	if c.Items[*item] == 0 {
		delete(c.Items, *item)
	}
}

func (c *Cart) RemoveProduct(item *product.Item) {
	delete(c.Items, *item)
}
