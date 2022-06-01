package product

import "testing"

func TestProductItem(t *testing.T) {
	t.Run("create iPad Pro", func(t *testing.T) {
		item := NewItem(IPadPro)

		if item.Name != "iPad Pro" {
			t.Errorf("Expected item name to be 'iPad Pro', but got %s", item.Name)
		}
	})

	t.Run("ipad pro with weight", func(t *testing.T) {
		item := NewItem(IPadPro)
		item.AddWeight(100)

		if item.Weight != 100 {
			t.Errorf("Expected item weight to be '100', but got %d", item.Weight)
		}
	})

	t.Run("create Apple Watch", func(t *testing.T) {
		item := NewItem(AppleWatch)
		item.AddWeight(50)

		if item.Name != "Apple Watch" {
			t.Errorf("Expected item name to be 'Apple Watch', but got %s", item.Name)
		}
	})

	t.Run("Apple Watch with weight", func(t *testing.T) {
		item := NewItem(IPadPro)
		item.AddWeight(50)

		if item.Weight != 50 {
			t.Errorf("Expected item weight to be '50', but got %d", item.Weight)
		}
	})

	t.Run("create Rice Cooker", func(t *testing.T) {
		item := NewItem(RiceCooker)
		item.AddWeight(2000)

		if item.Name != "Rice Cooker" {
			t.Errorf("Expected item name to be 'Rice Cooker', but got %s", item.Name)
		}
	})
}
