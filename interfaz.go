package mensajes

import "fmt"

type Formas interface {
	area() float64
	perimetro() float64
}

func Medidas(forma Formas) {
	fmt.Println(forma)
	fmt.Println("el area es de", forma.area())
	fmt.Println("el perimetro es de", forma.perimetro())
}
