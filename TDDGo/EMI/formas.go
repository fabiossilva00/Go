package emi

//Retangulo -
type Retangulo struct {
	Largura float64
	Altura  float64
}

//Circulo -
type Circulo struct {
	Raio float64
}

//Triangulo -
type Triangulo struct {
	Base   float64
	Altura float64
}

//Forma - Interface de Area
type Forma interface {
	Area() float64
}

//Perimetro -
func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

//Area - Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

//Area - Circulo
func (c Circulo) Area() float64 {
	return c.Raio * c.Raio * 3.14159
}

//Area - Triangulo
func (t Triangulo) Area() float64 {
	return t.Base * t.Altura / 3
}
