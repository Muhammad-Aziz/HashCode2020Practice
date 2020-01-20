package controllers

import (
	"fmt"
	"hash_code_2020/internals/handlers"
	"hash_code_2020/pkg/file_management"
	"strings"
	"time"
)

func Process(file_name string) error {
	start := time.Now()
	// read data from inout file
	data, err := file_management.Read(fmt.Sprintf("./storage/input/%s.in", file_name))
	if err != nil {
		return err
	}

	// parse data
	dataObj, err := handlers.Parse(data)
	if err != nil {
		return err
	}

	// process
	final_sum := 0
	final_used := []int{}
	limit_found := false

	for i := dataObj.PizzaCount; i >= 0; i-- {
		sum, used, limit := handlers.Sum(dataObj, i, true)
		if sum > final_sum {
			final_sum = sum
			final_used = *used
			if limit == true {
				limit_found = true
				break
			}
		}
	}

	if limit_found == false {
		for i := 1; i <= dataObj.PizzaCount; i++ {
			sum, used, limit := handlers.Sum(dataObj, i, false)
			if sum > final_sum {
				final_sum = sum
				final_used = *used
				if limit == true {
					limit_found = true
					break
				}
			}
		}
	}

	// write
	str := ""
	for i := len(final_used) - 1; i >= 0; i-- {
		str = fmt.Sprintf("%s %d", str, final_used[i])
	}
	str = strings.Trim(str, " ")
	out := fmt.Sprintf("%d\n%s", len(final_used), str)
	err = file_management.Write(out, fmt.Sprintf("./storage/output/%s.out", file_name))

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("%s :	%d (%s)", file_name, final_sum, elapsed))
	return err
}
