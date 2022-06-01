package pricing

import (
	"github.com/pallat/ddd/currency"
	"github.com/pallat/ddd/product"
)

var db = map[currency.Unit]map[product.ProductName]float64{
	currency.THB: {
		product.IPadPro: 37900.00,
	},
}

func Price(item *product.Item, currency currency.Unit) float64 {
	return db[currency][product.ProductName(item.Name)]
}
