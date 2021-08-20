package main

import (
	"fmt"
  "strings"
)


// func Reverse(s string) string {
//     runes := []rune(s)
//     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }

func minWindowDir(s string, t string) (string, int) {
	currentWindow := []string{}
	currentWindowLen := 0
	var patternMap map[string]int
	var currentPatternMapLen int
	var isStartWindowAvail bool

	tChars := strings.Split(t, "")
	sChars := strings.Split(s, "")
	minWindowYet := ""
	minWindowYetLen := 0
	buildPatternMap := func(t string) map[string]int {
		pMap := make(map[string]int)
		for _, val := range tChars {
			_, found := pMap[val]
			if found {
				pMap[val] += 1
			} else {
				pMap[val] = 1
			}
		}

		return pMap
	}

	realPatternMap := buildPatternMap(t)
	realPatternMapLen := len(realPatternMap)

	allocateResources := func() {
		isStartWindowAvail = false
		currentPatternMapLen = realPatternMapLen
		currentWindow = []string{}
		currentWindowLen = 0
		patternMap = make(map[string]int)
		for index, val := range realPatternMap {
			patternMap[index] = val
		}
	}

	allocateResources()

	for _, val := range sChars {
		_, found := patternMap[val]
		if found {
			if !isStartWindowAvail {
				isStartWindowAvail = true
			}
			patternMap[val] -= 1
			if patternMap[val] == 0 {
				delete(patternMap, val)
				currentPatternMapLen -= 1
			}
			currentWindow = append(currentWindow, val)
			currentWindowLen++
			if currentPatternMapLen == 0 {
        fmt.Println("current window:", currentWindow)
        fmt.Println("min window yet:", minWindowYet)

				if minWindowYetLen == 0 {
					minWindowYet = strings.Join(currentWindow, "")
					minWindowYetLen = currentWindowLen
				} else if minWindowYetLen > currentWindowLen {
					minWindowYet = strings.Join(currentWindow, "")
					minWindowYetLen = currentWindowLen
				}
				allocateResources()
			}
		} else {
			if isStartWindowAvail {
				currentWindow = append(currentWindow, val)
				currentWindowLen++
			} else {
				continue
			}
		}
	}

	return minWindowYet, minWindowYetLen
}

func minWindow(s string, t string) string {
	minWindowLeft, _ := minWindowDir(s, t)
  return minWindowLeft
	// if minWindowLeft == "" {
	// 	return minWindowLeft
	// }
  //
	// minWindowRight, minWindowRightLen := minWindowDir(Reverse(s), t)
	// if minWindowLeftLen > minWindowRightLen	 {
	// 	return Reverse(minWindowRight)
	// } else {
	// 	return minWindowLeft
	// }
}

func main() {
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
