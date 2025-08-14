package invoice

import "github.com/jcrispin/go-pdo/composicion/pkg/customer"

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
