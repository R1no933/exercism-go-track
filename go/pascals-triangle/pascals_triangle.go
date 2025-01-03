package pascal

func Triangle(n int) [][]int {
	res := [][]int{{1}}
	for len(res) < n {
		newline := []int{1}
		prevline := res[len(res)-1]
		for i := 1; i < len(res); i++ {
			newline = append(newline, prevline[i-1]+prevline[i])
		}
		newline = append(newline, 1)
		res = append(res, newline)
	}
	return res
}
