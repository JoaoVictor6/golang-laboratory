package main

import "math"

//interface
// São implicitas kkkkkkkkkkkkkkkkk mt foda se na minha interface tiver um método
//area, e receber um tipo que tiver esse msm método está tudo bem igualzin em typescript
type Forma interface {
	Area() float64
}

//Structure
type Retangulo struct {
	Altura  float64
	Largura float64
}

// r Retangulo equivale a um this, com ele podemos mexer no contexto em q essa structure foi chamada
func (r Retangulo) Area() float64 {
	return (r.Altura * r.Largura)
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {

}
