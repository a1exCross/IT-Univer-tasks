package main

import (
	"fmt"
	"os"
)

func main() {
	N := 0

	for N <= 0 || N > 100 {
		fmt.Println("\nВведите количество элементов массива")

		_, err := fmt.Fscan(os.Stdin, &N)
		if err != nil {
			fmt.Println(err)
		}
	}

	A := make([]int, N)

	fmt.Println("\nВведите элементы массива")

	for i := 0; i < N; i++ {
		for {
			_, err := fmt.Fscan(os.Stdin, &A[i])
			if err != nil {
				fmt.Println(err)
			}

			if A[i] < 1000 && A[i] >= -1000 {
				break
			}
		}
	}

	fmt.Println("\nВведенный массив:", A)

	K := 0

	for K <= 0 || K > 100 {
		fmt.Println("\nВведите количество сдвигов")

		_, err := fmt.Fscan(os.Stdin, &K)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("result =", solution(A, K))
}

func solution(A []int, K int) []int {
	shift_count := len(A) - K%len(A)
	A = append(A[shift_count:], A[:shift_count]...)

	return A
}
