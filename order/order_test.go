package order

import (
	"reflect"
	"testing"

	"github.com/pallat/ddd/cart"
	"github.com/pallat/ddd/product"
)

func TestCheckoutCart(t *testing.T) {
	t.Run("checkout empty cart", func(t *testing.T) {
		cart := cart.NewCart()

		order := NewOrder()
		po := order.Checkout(cart)

		if len(po.Items) != 0 {
			t.Errorf("Expected empty purchase order, but got %v", po)
		}
	})
	t.Run("checkout cart with 1 item", func(t *testing.T) {
		cart := cart.NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)

		order := NewOrder()
		po := order.Checkout(cart)

		expectedItem := product.NewItem(product.IPadPro)
		expected := &PurchaseOrder{
			Items: []PurchaseOrderItem{
				{
					Product:  *expectedItem,
					Quantity: 1,
					Price:    37900.00,
				},
			},
			SubTotal: 37900.00,
		}

		if !reflect.DeepEqual(po, expected) {
			t.Errorf("Expected %v, but got %v", expected, po)
		}
	})
	t.Run("checkout cart with 2 items", func(t *testing.T) {
		cart := cart.NewCart()
		cart.AddItem(product.NewItem(product.IPadPro), 1)
		cart.AddItem(product.NewItem(product.AppleWatch), 1)

		order := NewOrder()
		po := order.Checkout(cart)

		expectedItem1 := product.NewItem(product.IPadPro)
		expectedItem2 := product.NewItem(product.AppleWatch)
		expected := &PurchaseOrder{
			Items: []PurchaseOrderItem{
				{
					Product:  *expectedItem1,
					Quantity: 1,
					Price:    37900.00,
				},
				{
					Product:  *expectedItem2,
					Quantity: 1,
					Price:    14900.00,
				},
			},
			SubTotal: 52800.00,
		}

		if !reflect.DeepEqual(po, expected) {
			t.Errorf("Expected %v, but got %v", expected, po)
		}
	})
}

func TestTotalPrice(t *testing.T) {
	t.Run("total price include shipping cost 1 item", func(t *testing.T) {
		cart := cart.NewCart()
		item := product.NewItem(product.IPadPro)
		item.AddWeight(100)
		cart.AddItem(item, 1)

		order := NewOrder()
		po := order.Checkout(cart)

		expectedItem := product.NewItem(product.IPadPro)
		expectedItem.AddWeight(100)
		expected := &PurchaseOrder{
			Items: []PurchaseOrderItem{
				{
					Product:  *expectedItem,
					Quantity: 1,
					Price:    37900.00,
					Weight:   100,
				},
			},
			ShippingFee: 10.00,
			SubTotal:    37900.00,
			Total:       37910.00,
		}

		if !reflect.DeepEqual(po, expected) {
			t.Errorf("Expected %#v, but got %#v", expected, po)
		}
	})

	t.Run("total price include shipping cost 2 item", func(t *testing.T) {
		cart := cart.NewCart()
		item := product.NewItem(product.IPadPro)
		item.AddWeight(100)
		cart.AddItem(item, 1)

		item2 := product.NewItem(product.AppleWatch)
		item2.AddWeight(50)
		cart.AddItem(item2, 2)

		order := NewOrder()
		po := order.Checkout(cart)

		expectedItem := product.NewItem(product.IPadPro)
		expectedItem.AddWeight(100)

		expectedItem2 := product.NewItem(product.AppleWatch)
		expectedItem2.AddWeight(50)

		expected := &PurchaseOrder{
			Items: []PurchaseOrderItem{
				{
					Product:  *expectedItem,
					Quantity: 1,
					Price:    37900.00,
					Weight:   100,
				},
				{
					Product:  *expectedItem2,
					Quantity: 2,
					Price:    14900.00,
					Weight:   50,
				},
			},
			ShippingFee: 20.00,
			SubTotal:    67700.00,
			Total:       67720.00,
		}

		if !reflect.DeepEqual(po, expected) {
			t.Errorf("Expected %#v, but got %#v", expected, po)
		}
	})
}
