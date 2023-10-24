
func miniMaxSum(arr []int32) {
	// Write your code here
	sortedSlice := make([]int, len(arr))
	for i, v := range arr {
		sortedSlice[i] = int(v)
	}

	sort.Ints(sortedSlice)

	min := 0
	max := 0
	for i := 0; i < 4; i++ {
		min += sortedSlice[i]
		max += sortedSlice[len(sortedSlice)-i-1]
	}

	fmt.Printf("%v %v", min, max)
}