package main

func Numbers(start int, end int) []int {
	result := make([]int, end-start+1)
	for i, j := start, 0; i <= end; i, j = i+1, j+1 {
		result[j] = i
	}
	return result
}

func MapInt(items []int, projection func(item int) int) []int {
	result := make([]int, len(items))
	for i, v := range items {
		result[i] = projection(v)
	}
	return result
}

func Sum(items []int) int {
	sum := 0
	for _, x := range items {
		sum = sum + x
	}
	return sum
}

func Square(item int) int {
	return item * item
}
