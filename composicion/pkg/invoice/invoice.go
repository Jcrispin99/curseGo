package invoice

import (
	"github.com/jcrispin/go-pdo/composicion/pkg/customer"
	"github.com/jcrispin/go-pdo/composicion/pkg/invoiceitem"
)

// estructura de una factura
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// Nueva factura
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal es e; setter de total de la factura.
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
