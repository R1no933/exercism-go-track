package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	out := initial

	for _, vl := range s {
		out = fn(out, vl)
	}

	return out
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	out := initial
	size := len(s) - 1

	for i := size; i >= 0; i-- {
		out = fn(s[i], out)
	}

	return out
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := IntList([]int{})

	for _, vl := range s {
		if fn(vl) {
			out = append(out, vl)
		}
	}

	return out
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	out := IntList([]int{})

	for _, vl := range s {
		out = append(out, fn(vl))
	}

	return out
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, vl := range lists {
		s = append(s, vl...)
	}

	return s
}
