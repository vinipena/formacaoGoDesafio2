// Criar um código usando o operador % e for, onde exiba todos os numeros divisiveis por 3 entre 1 e 100.
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
