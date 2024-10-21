// Criar um programa que ao aparecer um numero multiplo de 3 imprimir "Pin" e ao aparecer um multiplo de 5 falar"Pan", nos numeros não citados devme imprimir o numero
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
