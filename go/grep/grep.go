package grep

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	flagsShort := map[byte]bool{}
	for _, f := range flags {
		flagsShort[f[len(f)-1]] = true
	}
	reString := ".*" + pattern + "./*"
	if _, ok := flagsShort['x']; ok {
		reString = "^" + pattern + "$"
	}
	if _, ok := flagsShort['i']; ok {
		reString = "(?i)" + pattern
	}
	_, vFlagOk := flagsShort['v']
	_, nFlagOk := flagsShort['n']
	_, lFlagOk := flagsShort['l']
	p := regexp.MustCompile(reString)
	fMatches, result := []string{}, []string{}
	for _, fName := range files {
		i, matches := 0, []string{}
		data, _ := os.ReadFile(fName)
		for _, line := range strings.Split(string(data), "\n") {
			if len(line) == 0 {
				continue
			}
			i += 1
			fLine := ""
			if len(files) > 1 {
				fLine += fName + ":"
			}
			if nFlagOk {
				fLine += fmt.Sprintf("%d:", i)
			}
			fLine += line
			matchRes := p.FindStringSubmatch(line)
			if len(matchRes) > 0 && !vFlagOk {
				matches = append(matches, fLine)
			} else if vFlagOk && len(matchRes) == 0 {
				matches = append(matches, fLine)
			}
		}
		if lFlagOk && len(matches) > 0 {
			fMatches = append(fMatches, fName)
		}
		result = append(result, matches...)
	}
	if lFlagOk && len(fMatches) > 0 {
		return fMatches
	}
	return result
}
