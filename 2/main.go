package main

import (
	"fmt"
)

func main() {
	/*
	   Примеры:
	   20
	   1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 9 9 10 10 10

	   7
	   9 3 9 3 9 7 9

	   15
	   9 7 4 2 8 9 7 4 1 5 2 1 5 8 3
	*/
	var N int

	for N <= 1 || N > 1000000 {
		fmt.Println("\nВведите количество элмементов массива")
		fmt.Scan(&N)
	}

	A := make([]int, N)

	fmt.Println("\nВведите элмементы массива")

	for i := 0; i < N; i++ {
		for {
			fmt.Scan(&A[i])

			if A[i] <= 1000000000 && A[i] >= 1 {
				break
			}
		}
	}

	fmt.Println("Ответ")
	Solution(A)
	//fmt.Println(Solution1(A))
}

func Solution(A []int) {
	hash_table := make(map[int]int, len(A))

	for i := 0; i < len(A); i++ {
		hash_table[A[i]] += 1
	}

	for key, v := range hash_table {
		if v <= 1 {
			fmt.Println(key)
		}
	}
}

/* func Solution1(A []int) int {
	for i := 0; i < len(A); i++ {
		count := 1
		for j := i + 1; j < len(A); j++ {
			if count > 1 {
				log.Println(j)
				break
			}

			if A[i] == A[j] {
				count++
			}
		}

		if count == 1 {
			return A[i]
		}
	}

	return 0
} */
