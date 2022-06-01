package product

import "testing"

func TestProductItem(t *testing.T) {
	t.Run("create iPad Pro", func(t *testing.T) {
		item := NewItem(IPadPro)

		if item.Name != "iPad Pro" {
			t.Errorf("Expected item name to be 'iPad Pro', but got %s", item.Name)
		}
	})
	t.Run("create Apple Watch", func(t *testing.T) {
		item := NewItem(AppleWatch)

		if item.Name != "Apple Watch" {
			t.Errorf("Expected item name to be 'Apple Watch', but got %s", item.Name)
		}
	})
	t.Run("create Rice Cooker", func(t *testing.T) {
		item := NewItem(RiceCooker)

		if item.Name != "Rice Cooker" {
			t.Errorf("Expected item name to be 'Rice Cooker', but got %s", item.Name)
		}
	})
}
