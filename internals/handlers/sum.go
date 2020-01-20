package handlers

import (
	"hash_code_2020/internals/models"
)

func Sum(dataObj *models.Data, ignore int, isPosition bool) (int, *[]int, bool) {
	sum := 0
	var used []int

	for i := dataObj.PizzaCount - 1; i >= 0; i-- {
		if isPosition == true && i == ignore {
			continue
		} else if isPosition == false && i >= dataObj.PizzaCount-ignore {
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
