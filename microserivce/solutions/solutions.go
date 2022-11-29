package solutions

import "sort"

func FindMissingElement(A []int) int {
	//судя по данным, массив всегда с 1, поэтому алгоритм упрощен
	sort.Ints(A)

	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}

	return -1
}

func SequenceCheck(A []int) int {
	sort.Ints(A)

	//судя по данным массив всегда идет с 1, поэтому алгоритм был значительно упрощен
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			return 0
		}
	}

	return 1
}

func WeirdArrayEntry(A []int) int {
	sort.Ints(A)

	prev := A[0]
	count := 1

	for i := 1; i < len(A); i++ {
		if A[i] == prev {
			prev = A[i]
			count++
		} else {
			if count%2 != 0 {
				return prev
			} else {
				count = 1
				prev = A[i]
			}
		}
	}

	return 0
}

func CyclicRotation(A []int, K int) []int {
	shift_count := len(A) - K%len(A)
	A = append(A[shift_count:], A[:shift_count]...)

	return A
}
