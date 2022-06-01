package pricing

import (
	"testing"

	"github.com/pallat/ddd/currency"
	"github.com/pallat/ddd/product"
)

func TestProductPrice(t *testing.T) {
	db = map[currency.Unit]map[product.ProductName]float64{
		currency.THB: {
			product.IPadPro: 37900.00,
		},
	}

	t.Run("create iPad Pro", func(t *testing.T) {
		item := product.NewItem(product.IPadPro)

		price := Price(item, currency.THB)

		if price != 37900 {
			t.Errorf("Expected price to be %f, but got %f\n", 37900.0, price)
		}
	})
}
