package order

import (
	"reflect"
	"testing"

	"github.com/pallat/ddd/cart"
	"github.com/pallat/ddd/product"
)

func TestCheckoutCart(t *testing.T) {
	order := NewOrder()

	cart := cart.NewCart()
	cart.AddItem(product.NewItem(product.IPadPro), 1)

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
		Total: 37900.00,
	}

	if !reflect.DeepEqual(po, expected) {
		t.Errorf("Expected %v, but got %v", expected, po)
	}
}
