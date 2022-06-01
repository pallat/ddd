package cart

import (
	"testing"

	"github.com/pallat/ddd/product"
)

func TestCartAddItem(t *testing.T) {
	t.Run("add iPad Pro into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 1 {
			t.Errorf("Expected cart to have 1 item, but got %d", sumqty)
		}
	})
	t.Run("add iPad Pro and Apple Watch into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 2 {
			t.Errorf("Expected cart to have 2 items, but got %d", sumqty)
		}
	})
	t.Run("Add iPad Pro, Apple Watch and 2 Rice Cooker into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)
		cart.AddItem(product.NewItem(product.RiceCooker), 2)

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 4 {
			t.Errorf("Expected cart to have 4 items, but got %d", sumqty)
		}
	})
}