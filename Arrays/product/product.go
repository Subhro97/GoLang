package product

import "github.com/Pallinder/go-randomdata"

type Product struct {
	id    int
	title string
	price float64
}

func New(title string, price float64) Product {
	return Product{
		id:    randomdata.Number(0, 50),
		title: title,
		price: price,
	}
}
