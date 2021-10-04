package main

func Perimetro(retangulo Retangulo) float64 {
	perimetro := 2 * (retangulo.Altura + retangulo.Largura)
	return perimetro
}
