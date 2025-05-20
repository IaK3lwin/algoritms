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

// function section soft algorithm
func SectionSoftAlgorithm(arr []int) []int {
	var NewSlice []int

	for range len(arr) {
		// fmt.Println(arr)
		smallesIndexValue := LookForMenor(arr)

		NewSlice = append(NewSlice, arr[smallesIndexValue])
		// adicionae slice[1, 5, 4, 7], ao 5 => [1, 5, 4, 7]

		arr = append(arr[:smallesIndexValue], arr[smallesIndexValue+1:]...)

	}

	return NewSlice
}

func main() {
	fmt.Println("creating a section sort algorithm  ")

	sliceteste := []int{1, 5, 4, 7, 0, 5}

	fmt.Println(LookForMenor(sliceteste))

	sliceSoft := SectionSoftAlgorithm(sliceteste)

	fmt.Println(sliceSoft)

}
