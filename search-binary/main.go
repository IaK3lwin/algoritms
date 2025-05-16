package main

import (
	"errors"
	"fmt"
)

//binary search

var lista []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Pesquisa_binaria(list []int, item int) (int, error) {
	//definir o nosso \"açvp"
	//onde a gente estará buscando pelo noe
	baixo := 0
	alto := len(list) - 1

	for baixo <= alto {
		//intermédiário do nosso array
		meio := (alto + baixo) / 2

		//chute
		chute := list[meio]

		if chute == item {
			return meio, nil
		}
		// 5   >  7
		if chute < item {
			baixo = meio + 1
		} else {
			alto = meio - 1
		}
	}

	return 0, errors.New("valor não encontrado")

}

func main() {
	fmt.Println("usando pesquisa binária ")
	fmt.Println(Pesquisa_binaria(lista, 7))

}
