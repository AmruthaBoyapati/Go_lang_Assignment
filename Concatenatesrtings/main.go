package main

import (
	"fmt"
)

func smallestConcat(arr []string) string {
	minStr := ""
	first := true

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				concat1 := arr[i] + arr[j]
				concat2 := arr[j] + arr[i]

				// Take lexicographically smaller of the two
				candidate := concat1
				if concat2 < concat1 {
					candidate = concat2
				}

				if first || candidate < minStr {
					minStr = candidate
					first = false
				}
			}
		}
	}

	return minStr
}

func main() {
	arr := []string{"aab", "bcddbc", "aa", "aazef"}
	fmt.Println(smallestConcat(arr)) 
}
