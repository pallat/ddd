package cart

import (
	"fmt"

	"github.com/pallat/ddd/currency"
	"github.com/pallat/ddd/pricing"
	"github.com/pallat/ddd/product"
)

type Cart struct {
	Items        map[product.Item]int64
	IsCheckedOut bool
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
	defer fmt.Printf("%s has removed\n", item.Name)
	delete(c.Items, *item)
}

func (c *Cart) Checkout() float64 {
	var total float64

	for item, quantity := range c.Items {
		total += pricing.Price(&item, currency.THB) * float64(quantity)
	}

	c.IsCheckedOut = true
	return total
}
