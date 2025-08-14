package course

import (
	"fmt"
	"strings"
)

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func NewCourse(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}

func (c *course) PrintCourse() {
	// Usar strings.Builder es mucho más eficiente para construir strings.
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Las clases del curso %s son:\n", c.Name))
	for _, class := range c.Classes {
		builder.WriteString("- " + class + "\n")
	}
	fmt.Print(builder.String())
}
