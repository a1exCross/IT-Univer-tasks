package main

import (
	"fmt"
)

func main() {

	var N int

	for N <= 1 || N > 100000 {
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

	fmt.Println("Ответ: ")
	res := Solution(A)

	if res == 1 {
		fmt.Println("Это последовательность")
	} else if res == 0 {
		fmt.Println("Это не последовательность")
	}

}

func Solution(A []int) int {
	hash_table := make(map[int]int, len(A))

	min := A[0]

	for i := 0; i < len(A); i++ {
		hash_table[A[i]] += 1

		if hash_table[A[i]] > 1 { //проверка на дублирующиеся числа
			return 0
		}

		if A[i] < min {
			min = A[i] // поиск минимального для пострения последовательности
		}
	}

	hash_table = make(map[int]int, len(A)) // очистка хэш-мапа

	for i := 0; i < len(A); i++ {
		hash_table[min] = 1 //помечаем последовательность единицей
		min++
	}

	for i := 0; i < len(A); i++ {
		if hash_table[A[i]] != 1 { // если в таблице нет элемента из массива, то это не последовательность
			return 0
		}
	}

	return 1
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
