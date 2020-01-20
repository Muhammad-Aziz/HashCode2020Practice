package handlers

import (
	"errors"
	"hash_code_2020/internals/models"
	"strconv"
	"strings"
)

func Parse(data string) (*models.Data, error) {
	dataObj := models.Data{}
	var err error

	lines := strings.Split(data, "\n")

	// parse first line (guest count + number of pizza types)
	words := strings.Split(lines[0], " ")
	dataObj.GuestCount, err = strconv.Atoi(words[0])
	if err != nil {
		return &dataObj, err
	}
	dataObj.PizzaCount, err = strconv.Atoi(words[1])
	if err != nil {
		return &dataObj, err
	}

	words = strings.Split(lines[1], " ")
	for i := range words {
		if words[i] != "" {
			sliceCount, err := strconv.Atoi(words[i])
			if err != nil {
				return &dataObj, err
			}
			p := models.Pizza{
				Slices:   sliceCount,
				Position: i,
			}
			//p.Remaining = m.GuestCount - p.Slices

			dataObj.Slices = append(dataObj.Slices, p)
		}
	}

	if len(dataObj.Slices) != dataObj.PizzaCount {
		err = errors.New("Something went wrong")
	}
	return &dataObj, err
}
