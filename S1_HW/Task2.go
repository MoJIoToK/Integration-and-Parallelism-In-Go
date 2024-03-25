package main

import (
	"fmt"
	"math/rand"
)

const sizeArray = 6

func main() {
	var array [sizeArray]int
	for i := range array {
		array[i] = rand.Intn(20) - 10 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println("Source array: ", array)

	array = bubbleSort(array)
	fmt.Println("Sorted array: ", array)

}

func bubbleSort(array [sizeArray]int) [sizeArray]int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array

}
