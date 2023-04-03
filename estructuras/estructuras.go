package estructuras

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p *Product) Save(products []Product) {
	products = append(products, *p)
}

func (p *Product) GetAll(products []Product) {
	fmt.Println(products)
}

func GetById(products []Product, id int) Product {
	for _, product := range products {
		if product.Id == id {
			return product
		}
	}
	return products[0]
}
