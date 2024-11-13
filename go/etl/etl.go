package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newDataForm := make(map[string]int)

	for scr, wds := range in {
		for _, wd := range wds {
			newDataForm[strings.ToLower(wd)] = scr
		}
	}

	return newDataForm
}
