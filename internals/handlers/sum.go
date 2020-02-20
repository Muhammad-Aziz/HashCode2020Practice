package handlers

import (
	"hash_code_2020/practice/internals/models"
)

func Sum(dataObj *models.Data, ignore int) (int, *[]int, bool) {
	sum := 0
	var used []int

	skip := dataObj.Slices[dataObj.PizzaCount-ignore].Slices

	for i := dataObj.PizzaCount - 1; i >= 0; i-- {
		if skip == dataObj.Slices[i].Slices {
			continue
		}
		sum = sum + dataObj.Slices[i].Slices
		if sum <= dataObj.GuestCount {
			used = append(used, dataObj.Slices[i].Position)
		} else {
			sum = sum - dataObj.Slices[i].Slices
		}
	}

	return sum, &used, sum == dataObj.GuestCount
}
