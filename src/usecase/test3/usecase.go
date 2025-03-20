package test3

func FindOddInt(input []int) int {

	nums := make(map[int]int)

	for _, num := range input {
		nums[num]++
    }

	for i, n := range nums {
		if n%2 != 0 {
			return i
		}
	}

	return 0
}