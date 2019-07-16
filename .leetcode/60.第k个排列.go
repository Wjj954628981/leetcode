import "strconv"

func getPermutation(n int, k int) string {
	var ret string
	nums := make(map[int]int)
	var FAC = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	var cantor func(i int, x int)
	cantor = func(i int, x int) {
		if i == -1 {
			return
		}
		quotient := int(x / FAC[i])
		remainder := x % FAC[i]
		var index = 0
		var j = 1
		for ; j <= n; j++ {
			if _, ok := nums[j]; ok {
				continue
			}
			index++
			if index > quotient {
				break
			}
		}
		nums[j] = j
		ret = ret + strconv.Itoa(j)
		i--
		cantor(i, remainder)
	}
	cantor(n-1, k-1)
	return ret
}