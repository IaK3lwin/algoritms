package main

import "fmt"

// function look for a menor
func LookForMenor(arr []int) int {
	//get small item
	smallestItem := arr[0]
	// and small index
	smallesIndex := 0

	//loop in range values
	for index, _ := range arr {
		//se element in index current is less then smallestItem
		if arr[index] < smallestItem {

			smallestItem = arr[index]
			smallesIndex = index
		}

	}
	//retorno i indez do item menor
	return smallesIndex
}





func main() {
	fmt.Println("creating a section sort algorithm  ")
	arrayteste := []int{1, 5, 4, 7, 0, 5}

	fmt.Println(LookForMenor(arrayteste))

}
