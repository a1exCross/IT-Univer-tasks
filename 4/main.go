package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int

	for N <= 0 || N > 100000 {
		fmt.Println("\nВведите количество элмементов массива")
		fmt.Scan(&N)
	}

	A := make([]int, N)

	fmt.Println("\nВведите элмементы массива")

	for i := 0; i < N; i++ {
		for {
			fmt.Scan(&A[i])

			if A[i] <= N+1 && A[i] >= 1 {
				break
			}
		}
	}

	fmt.Println("Ответ: ")
	res := Solution(A, N)

	if res == -1 {
		fmt.Println("Пропущенные числа отсутствуют")
	} else {
		fmt.Println("Проупщенное число", res)
	}

}

func Solution(A []int, N int) int {
	/* hash_table := make(map[int]int, N+1)

	min := A[0]

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i] // поиск минимального для пострения последовательности
		}
	}

	arr_sort := make([]int, N)

	for i := 0; i < len(A); i++ {
		arr_sort[i] = min
		hash_table[min] = 1
		min++
	}

	i := 0

	for key, _ := range hash_table {
		if hash_table[A[i]] == 0 {
			return key
		}

		i++
	}

	return 1 */

	/* for i := 0; i < len(A); i++ {
		if hash_table[A[i]] == 0 {
			return hash_table[i]
		}
	} */

	sort.Ints(A)

	min := A[0]

	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i] // поиск минимального для пострения последовательности
		}
	}

	arr_sort := make([]int, N)

	for i := 0; i < len(A); i++ {
		arr_sort[i] = min
		min++
	}

	for i := 0; i < len(A); i++ {
		if A[i] != arr_sort[i] {
			return arr_sort[i]
		}
	}

	return -1
}

/* func Solution(A []int) int {
	hash_table := make(map[int]int, len(A))

	for i := 0; i < len(A); i++ {
		hash_table[A[i]] += 1
	}

	var res = 0

	for key, v := range hash_table {
		if v > 1 {
			delete(hash_table, key)
			continue
		} else {
			return key
		}
	}

	return res
}

func Solution1(A []int) int {

	for i := 0; i < len(A); i++ {
		count := 0
		for j := 1; j < len(A); j++ {
			if count > 1 {
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
}
*/
