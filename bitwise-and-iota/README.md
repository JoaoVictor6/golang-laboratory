## IOTA
Gera valores iniciais para constantes

```go
const (
	a = iota*10 // vai adicionar IOTA automaticamente nos de baixo e multiplicar os valores por 10
	b
	c
)

const (
	_ = iota // suspendendo alguns valores para começarmos do 1 ao invés do zero
	e = iota
	f = iota
)

func main(){
	fmt.Println(a,b,c) // 0, 10, 20
	fmt.Println(e,f) // 1, 2
}
```
### Bibliografia
[Lógica binária](https://pt.wikipedia.org/wiki/L%C3%B3gica_bin%C3%A1ria)

[GO - Bitwise operators](https://www.tutorialspoint.com/go/go_bitwise_operators.htm)

[Conheça os operadores Bitwise](https://imasters.com.br/desenvolvimento/conheca-os-operadores-bitwise-bit-bit)
