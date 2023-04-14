package mensajes

type Circulo struct {
	Radio float64
}

func (cir *Circulo) area() float64 {
	return 3.14 * (cir.Radio * cir.Radio)
}

func (cir *Circulo) perimetro() float64 {
	return 2 * 3.14 * cir.Radio
}
