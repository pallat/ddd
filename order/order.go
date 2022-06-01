package order

import (
	"github.com/pallat/ddd/cart"
	"github.com/pallat/ddd/currency"
	"github.com/pallat/ddd/pricing"
	"github.com/pallat/ddd/product"
)

type PurchaseOrder struct {
	Items []PurchaseOrderItem
	Total float64
}

type PurchaseOrderItem struct {
	Product  product.Item
	Quantity int64
	Price    float64
}

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) Checkout(cart *cart.Cart) *PurchaseOrder {
	po := &PurchaseOrder{}

	for item, quantity := range cart.Items {
		po.Items = append(po.Items, PurchaseOrderItem{
			Product:  item,
			Quantity: quantity,
			Price:    pricing.Price(&item, currency.THB),
		})
	}

	po.Total = cart.Checkout()

	return po
}
