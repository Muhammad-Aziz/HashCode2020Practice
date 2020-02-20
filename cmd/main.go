package main

import "hash_code_2020/practice/internals/controllers"

func main() {
	// set file names to be used
	input_files := []string{"a_example", "b_small", "c_medium", "d_quite_big", "e_also_big"}

	// iterate over input files for processing
	for i := 0; i < 5; i++ {
		check(controllers.Process(input_files[i]))
	}
}

// panic on error
// since we are expecting to finish without errors
// it's better to stop than get wrong results :)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
