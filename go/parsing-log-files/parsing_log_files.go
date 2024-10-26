package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	rg := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return rg.MatchString(text)
}

func SplitLogLine(text string) []string {
	rg := regexp.MustCompile(`<[*~=-]*>`)
	return rg.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	rg := regexp.MustCompile(`(?i)".*password.*"`)
	res := 0
	for _, line := range lines {
		if rg.MatchString(line) {
			res++
		}
	}
	return res
}

func RemoveEndOfLineText(text string) string {
	rg := regexp.MustCompile(`end-of-line\d*`)
	return rg.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	rg := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, line := range lines {
		founds := rg.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}
