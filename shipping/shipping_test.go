package shipping

import (
	"testing"

	"github.com/pallat/ddd/currency"
)

func TestCalculateShippingPrice(t *testing.T) {

	t.Run("calculate item weight 100g", func(t *testing.T) {
		input := 100
		expected := 10.0

		shippingPrice := Price(int64(input), currency.THB)

		if shippingPrice != expected {
			t.Errorf("Expected shipping price to be %f, but got %f\n", 10.0, shippingPrice)
		}
	})
}
