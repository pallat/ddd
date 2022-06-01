package shipping

import "github.com/pallat/ddd/currency"

var db = map[currency.Unit]float64{
	currency.THB: 0.1,
}

func Fee(weight int64, currency currency.Unit) float64 {
	return float64(weight) * db[currency]
}
