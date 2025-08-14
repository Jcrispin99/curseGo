package main

import (
	"fmt"

	"github.com/jcrispin/go-pdo/estructuras/course"
)

func main() {
	Go := course.NewCourse("Go desde cero", 0, true)
	Go.SetUserIDs([]uint{1, 2, 3, 4, 5})
	Go.SetClasses(map[uint]string{
		1: "Introducción1",
		2: "Tipos de datos",
		3: "Estructuras",
		4: "Funciones",
		5: "Punteros",
	})
	// Go.Classes = map[uint]string{
	// 	1: "Introducción",
	// 	2: "Tipos de datos",
	// 	3: "Estructuras",
	// 	4: "Funciones",
	// 	5: "Punteros",
	// }
	// Go := &course.Course{
	// 	Name:    "Go desde cero",
	// 	Price:   12,
	// 	IsFree:  true,
	// 	UserIDs: []uint{1, 2, 3, 4, 5},
	// 	Classes: map[uint]string{
	// 		1: "Introducción",
	// 		2: "Tipos de datos",
	// 		3: "Estructuras",
	// 		4: "Funciones",
	// 		5: "Punteros",
	// 	},
	// }

	Go.SetName("PDO con Go")
	Go.PrintCourse()
	fmt.Println("nombre:", Go.Name())
}
