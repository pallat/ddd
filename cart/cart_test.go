package cart

import (
	"testing"

	"github.com/pallat/ddd/product"
)

func TestCartAddItem(t *testing.T) {
	t.Run("add iPad Pro into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro))

		if len(cart.Items) != 1 {
			t.Errorf("Expected cart to have 1 item, but got %d", len(cart.Items))
		}
	})
	t.Run("add iPad Pro and Apple Watch into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro))
		cart.AddItem(product.NewItem(product.AppleWatch))

		if len(cart.Items) != 2 {
			t.Errorf("Expected cart to have 2 items, but got %d", len(cart.Items))
		}
	})
	t.Run("Add iPad Pro, Apple Watch and 2 Rice Cooker into Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro))
		cart.AddItem(product.NewItem(product.AppleWatch))
		cart.AddItem(product.NewItem(product.RiceCooker))
		cart.AddItem(product.NewItem(product.RiceCooker))

		if len(cart.Items) != 4 {
			t.Errorf("Expected cart to have 4 items, but got %d", len(cart.Items))
		}
	})
}
