package main

import "fmt"

// Функция для пересечения двух множеств
func intersectSets(set1, set2 []int) []int {
	// Используем карту для хранения элементов первого множества
	elementMap := make(map[int]bool)
	for _, num := range set1 {
		elementMap[num] = true
	}

	// Ищем пересечения
	var intersection []int
	for _, num := range set2 {
		//проверяем существует ли ключ в map и имеет ли он значение true
		if elementMap[num] {
			intersection = append(intersection, num)

			delete(elementMap, num) // Удаляем, чтобы избежать дублирования
		}
	}

	return intersection
}

func main() {
	// Примеры двух неупорядоченных множеств
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Вычисляем пересечение
	result := intersectSets(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", result)
}
