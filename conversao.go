/* Converter o ponto de
ebulição da água  de kelvin para Celsis

C = K - 273

*/

package main

import "fmt"

const ebulicaoAguaK float64 = 373.0

func main() {
	tempAgua := ebulicaoAguaK
	tempC := tempAgua - 273

	fmt.Printf("Ponto de bulição da água em K = %g", tempAgua)
	fmt.Printf("\nPonto de ebulição da água convertido de K para celsius C = %g", tempC)

}
