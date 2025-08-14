package main

import (
	"fmt"

	"github.com/jcrispin/go-pdo/composicion/pkg/customer"
	"github.com/jcrispin/go-pdo/composicion/pkg/invoice"
	"github.com/jcrispin/go-pdo/composicion/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Spain",
		"Barcelona",
		customer.New("John Doe", "123 Main St", "555-1234"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Widget A", 19.99),
			invoiceitem.New(2, "Widget B", 29.99),
			invoiceitem.New(3, "Widget C", 9.99),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
