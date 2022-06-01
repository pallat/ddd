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

func TestCartRemoveItem(t *testing.T) {
	t.Run("remove iPad Pro from Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)
		cart.AddItem(product.NewItem(product.RiceCooker), 2)

		cart.RemoveItem(product.NewItem(product.IPadPro), 1)

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 3 {
			t.Errorf("Expected cart to have 3 items, but got %d\n%#v\n", sumqty, cart.Items)
		}
	})

	t.Run("remove 1 Rice Cooker from Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)
		cart.AddItem(product.NewItem(product.RiceCooker), 2)

		cart.RemoveItem(product.NewItem(product.RiceCooker), 1)

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 3 {
			t.Errorf("Expected cart to have 3 items, but got %d\n%#v\n", sumqty, cart.Items)
		}
	})
}

func TestCartRemoveProduct(t *testing.T) {
	t.Run("remove iPad Pro from Cart", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)
		cart.AddItem(product.NewItem(product.RiceCooker), 2)

		cart.RemoveProduct(product.NewItem(product.IPadPro))

		var sumqty int64
		for _, qty := range cart.Items {
			sumqty += qty
		}

		if sumqty != 3 {
			t.Errorf("Expected cart to have 3 items, but got %d\n%#v\n", sumqty, cart.Items)
		}
	})
}

func TestCartCheckout(t *testing.T) {
	t.Run("checkout with empty cart", func(t *testing.T) {
		cart := NewCart()
		total := cart.Checkout()

		if total != 0 {
			t.Errorf("Expected total to be 0, but got %f", total)
		}
	})
	t.Run("checkout with 1 iPad Pro", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)

		total := cart.Checkout()

		if total != 37900 {
			t.Errorf("Expected total to be 37900, but got %f", total)
		}
	})
	t.Run("checkout set true to IsCheckedOut when checkout has called", func(t *testing.T) {
		cart := NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)

		_ = cart.Checkout()

		if cart.IsCheckedOut != true {
			t.Errorf("Expected IsCheckedOut to be true, but got %t", cart.IsCheckedOut)
		}
	})
}
