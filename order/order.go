package order

import (
	"github.com/pallat/ddd/cart"
	"github.com/pallat/ddd/currency"
	"github.com/pallat/ddd/pricing"
	"github.com/pallat/ddd/product"
	"github.com/pallat/ddd/shipping"
)

type PurchaseOrder struct {
	Items       []PurchaseOrderItem
	SubTotal    float64
	ShippingFee float64
	Total       float64
}

type PurchaseOrderItem struct {
	Product  product.Item
	Quantity int64
	Price    float64
	Weight   int64
}

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) Checkout(cart *cart.Cart) *PurchaseOrder {
	var sumWeight int64
	po := &PurchaseOrder{}

	for item, quantity := range cart.Items {
		po.Items = append(po.Items, PurchaseOrderItem{
			Product:  item,
			Quantity: quantity,
			Price:    pricing.Price(&item, currency.THB),
			Weight:   item.Weight,
		})

		sumWeight += item.Weight * quantity
	}

<<<<<<< HEAD
	po.SubTotal = cart.Checkout()
	po.ShippingFee = shipping.Price(sumWeight, currency.THB)
=======
	po.SubTotal = subtotal
	po.ShippingFee = shipping.Fee(sumWeight, currency.THB)
>>>>>>> 099d420 (Rename function)
	po.Total = po.SubTotal + po.ShippingFee

	return po
}
