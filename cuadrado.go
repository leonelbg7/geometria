package mensajes

type Cuadrado struct {
	Lado float64
}

func (cua *Cuadrado) perimetro() float64 {
	return (2 * cua.Lado) + (2 * cua.Lado)
}

func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}
