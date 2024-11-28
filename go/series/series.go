package series

func All(n int, s string) []string {
	var res []string
	start := 0
	end := start + n

	for end <= len(s) {
		res = append(res, s[start:end])
		start += 1
		end += 1
	}

	return res

}

func UnsafeFirst(n int, s string) string {
	first := ""
	for i := 0; i < n; i++ {
		first += string(s[i])
	}

	return first
}
