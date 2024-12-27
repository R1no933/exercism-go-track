package allergies

var algyMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(allergies uint) []string {
	alist := []string{}
	for val, label := range algyMap {
		if val&allergies > 0 {
			alist = append(alist, label)
		}
	}
	return alist
}
func AllergicTo(allergies uint, allergen string) bool {
	for val, label := range algyMap {
		if label == allergen && val&allergies > 0 {
			return true
		}
	}
	return false
}
