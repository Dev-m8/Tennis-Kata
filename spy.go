package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strings"
)
func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i,v := range t {
		need[v] = need[v] + 1
		
	}
	size, total := len(s), len(t)
	min := size + 1
	res := ""
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}
		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}

	}

	return res
}

// function to ignore characters that are not alphabets
func isLetter(c rune) bool { 
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func verify(message, magazine string) bool {
	mapOfLettersToFrequency := make(map[string]int) 
	indexUpdate := 0 
	result := true
	if len(message) < 1 || len(message) > len(magazine) {
		result = false
	} else {
		for _, v := range message {
			if result == false {
				break
			}
			if isLetter(v) == false { 
				continue
			}
			messageChar := strings.ToLower(string(v)) 
			if value, exists := mapOfLettersToFrequency[messageChar]; exists && value != 0 {
				mapOfLettersToFrequency[messageChar]--
				result = true
			} else {
				for i := indexUpdate; i < len(magazine); i++ { 
					magazineChar := strings.ToLower(string(magazine[i])) 
					if messageChar == magazineChar {
						result = true
						indexUpdate = i + 1 
						break
					} else {
						mapOfLettersToFrequency[magazineChar]++ 
						result = false
					}
				}

			}

		}

	}
	return result
}

func main() {

	S := "ADOBECODEBANC"
    T := "ABC"

	fmt.Println(minWindow(S,T))

}


