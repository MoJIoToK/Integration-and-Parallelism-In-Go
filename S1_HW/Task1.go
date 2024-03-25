package main

import "fmt"

const (
	sizeA      = 4
	sizeB      = 5
	sizeResult = sizeA + sizeB
)

func main() {
	arrayA := [sizeA]int{1, 3, 5, 7}
	arrayB := [sizeB]int{2, 4, 6, 8, 9}

	mergeArray := mergeArray(arrayA, arrayB)
	fmt.Println(mergeArray)

	var n [len(arrayA) + len(arrayB)]int
	copy(n[:], arrayA[:])
	copy(n[len(arrayA):], arrayB[:])
	fmt.Printf("%T", n)

}

func mergeArray(a [sizeA]int, b [sizeB]int) [sizeResult]int {
	var result [sizeResult]int
	var i, j, k = 0, 0, 0

	for k < (len(a) + len(b)) {
		if i < len(a) {
			if a[i] < b[j] {
				result[k] = a[i]
				i++
			} else {
				result[k] = b[j]
				j++
			}
			k++
		} else {
			result[k] = b[j]
			j++
			k++
		}
	}

	//for j < len(b) {
	//	result[k] = b[j]
	//	j++
	//	k++
	//}

	//for i < len(a) && j < len(b) {
	//	if a[i] < b[j] {
	//		result[k] = a[i]
	//		i++
	//	} else {
	//		result[k] = b[j]
	//		j++
	//	}
	//	k++
	//}
	//
	//for j < len(b) {
	//	result[k] = b[j]
	//	j++
	//	k++
	//}

	return result
}
