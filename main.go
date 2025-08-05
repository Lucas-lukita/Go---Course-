package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var array1 [5]int
	array1[0] = 10
	fmt.Println(array1)

	// declarando os arrays com 5 posições de strings

	array2 := [4]string{"Vila nova", "Fluminense", "Goias", "Atletico"}
	fmt.Println(array2)

	// aqui vamos passar um modelo de array que vai amazenar o tamanho do array conforme os valores que forem inseridos
	array4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array4)
}
