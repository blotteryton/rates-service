package currency

import "golang.org/x/exp/slices"

type Currency struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewCurrency(id uint32, name string, code string) *Currency {
	currency := Currency{}
	currency.ID = id
	currency.Name = name
	currency.Code = code
	return &currency
}

func Fake() []Currency {
	return []Currency{
		*NewCurrency(1, "Российский рубль", "RUB"),
		*NewCurrency(2, "Американский доллар", "USD"),
	}
}

func Find(code string) *Currency {
	currencies := Fake()
	idx := slices.IndexFunc(currencies, func(c Currency) bool {
		return c.Code == code
	})

	if idx == -1 {
		return nil
	} else {
		return &currencies[idx]
	}
}
