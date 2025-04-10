Concatenação

package main
import "fmt"
func main() {
var n1, n2, n3 float64
var multi, multi2, multi3, soma, quadrado float64
fmt.Print("Digite o primeiro número: ")
fmt.Scan(&n1)
fmt.Print("Digite o segundo número: ")
fmt.Scan(&n2)
fmt.Print("Digite o terceiro número: ")
fmt.Scan(&n3)
multi = n1 * 100
multi2 = n2 * 10
multi3 = n3 * 1
soma = multi + multi2 + multi3
quadrado = soma * soma
if (n1n1 < 100) && (n2n2 < 100) && (n3*n3 < 100) {
fmt.Printf("A concatenação será: %.0f\n", soma)
fmt.Printf("O quadrado dela será: %.0f\n", quadrado)
} else {
fmt.Println("DÍGITO INVALIDO")
}
}
