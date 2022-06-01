package product

type ProductType int

type ProductName string

const (
	IPadPro    ProductName = "iPad Pro"
	AppleWatch             = "Apple Watch"
	RiceCooker             = "Rice Cooker"
)

type Item struct {
	Name   string
	Weight int64
}

func NewItem(productName ProductName) *Item {
	return &Item{Name: string(productName)}
}

func (i *Item) AddWeight(weight int64) {
	i.Weight = weight
}
