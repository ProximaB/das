package util

func InterfaceSliceToIntSlice(input []interface{}) []int {
	output := make([]int, 0)
	for _, each := range input {
		output = append(output, int(each.(int)))
	}
	return output
}

func InterfaceMapToStringMap(input map[string]interface{}) map[string]bool {
	output := make(map[string]bool)
	for k, v := range input {
		output[k] = v.(bool)
	}
	return output
}
