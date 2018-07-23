package util

func InterfaceSliceToIntSlice(input []interface{}) []int {
	output := make([]int, 0)
	for _, each := range input {
		output = append(output, each.(int))
	}
	return output
}
