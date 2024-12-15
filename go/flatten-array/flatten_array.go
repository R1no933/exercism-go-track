package flatten

func Flatten(nested interface{}) []interface{} {
	res := make([]interface{}, 0)
	list, isSlice := nested.([]interface{})
	if !isSlice {
		return append(res, nested)
	}
	for _, sub := range list {
		if sub != nil {
			res = append(res, Flatten(sub)...)
		}
	}
	return res
}
