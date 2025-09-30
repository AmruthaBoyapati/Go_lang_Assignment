package main

import "fmt"

func removeDuplicates(arr []int) []int {
	result := []int{}

	for i := 0; i < len(arr); i++ {
		duplicate := false
		// check if arr[i] is already in result
		for j := 0; j < len(result); j++ {
			if arr[i] == result[j] {
				duplicate = true
				break
			}
		}
		// if not duplicate, add to result
		if !duplicate {
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 2, 5, 4}
	fmt.Println(removeDuplicates(arr)) 
}
