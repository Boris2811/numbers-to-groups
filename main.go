package main

import (
	"fmt"
	"time"
)

type groupsType [][]int

func main() {
	var groups groupsType
	firstGroup := []int{1}
	secondGroup := []int{2}
	groups = append(groups, firstGroup)
	groups = append(groups, secondGroup)
	var inputNumber int
	groupsCount := 1
	fmt.Print("Введите число: ")
	fmt.Scanf("%d\n", &inputNumber)

	startTime := time.Now()
	for counter := 3; counter <= inputNumber; counter++ {
		numberAdded := false
		for groupsCounter := 0; groupsCounter < groupsCount; groupsCounter++ {
			canAdd := canAddToGroup(&groups[groupsCounter], counter)
			if canAdd {
				groups[groupsCounter] = append(groups[groupsCounter], counter)
				numberAdded = true
				break
			}
		}
		if !numberAdded {
			groups = append(groups, []int{counter})
			groupsCount++
		}
	}
	stopTime := time.Now()

	fmt.Println("Количество групп:", len(groups))

	var countInfoGroups int
	if len(groups) < 10 {
		countInfoGroups = len(groups)
	} else {
		countInfoGroups = 10
	}

	for i := 0; i < countInfoGroups; i++ {
		fmt.Println(i+1, groups[i])
	}

	fmt.Printf("Время затраченное на подсчёт: %v\n", stopTime.Sub(startTime))
}

func canAddToGroup(group *[]int, number int) bool {
	elemCount := len(*group)
	for counter := 0; counter < elemCount; counter++ {
		if number%(*group)[counter] == 0 {
			return false
		}
	}
	return true
}
