package product

type ProductType int

const (
	IPadPro ProductType = iota
	AppleWatch
	RiceCooker
)

type Item struct {
	Name string
}

func NewItem(productType ProductType) *Item {
	switch productType {
	case IPadPro:
		return &Item{Name: "iPad Pro"}
	case AppleWatch:
		return &Item{Name: "Apple Watch"}
	case RiceCooker:
		return &Item{Name: "Rice Cooker"}
	default:
		return nil
	}
}
