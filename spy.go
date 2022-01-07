package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}
	for i := range t {
		need[t[i]]++
	}

	size, total := len(s), len(t)

	min := size + 1
	res := ""

	// s[i:j+1] 就是 window
	// count 用于统计已有的 t 中字母的数量。
	// count == total 表示已经收集完需要的全部字母
	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// 出现了 window 中缺失的字母
			count++
		}
		have[s[j]]++

		// 保证 window 不丢失所需字母的前提下
		// 让 i 尽可能的大
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

// function validates if letters in message are contained in magazine i.e. message ⊆ magazine
func verify(message, magazine string) bool {
	mapOfLettersToFrequency := make(map[string]int) // build a map of letters seen when searching magazine
	indexUpdate := 0 //increment index of last char searched in magazine and store in this var
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
			messageChar := strings.ToLower(string(v)) //convert message char to lower case as we will ignore case sensitivity when matching letters
			if value, exists := mapOfLettersToFrequency[messageChar]; exists && value != 0 {
				mapOfLettersToFrequency[messageChar]--
				result = true
			} else {
				for i := indexUpdate; i < len(magazine); i++ { //iterate over magazine to find matching letters
					magazineChar := strings.ToLower(string(magazine[i])) //convert magazine char to lower case as we will ignore case sensitivity when matching letters
					if messageChar == magazineChar {
						result = true
						indexUpdate = i + 1 //increment indexUpdate by 1 so as to avoid searching previous indexes of magazine
						break
					} else {
						mapOfLettersToFrequency[magazineChar]++ //build map of characters and increment value 
						result = false
					}
				}

			}

		}

	}
	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	message, _ := reader.ReadString('\n')

	fmt.Println("Enter magazine: ")
	magazine, _ := reader.ReadString('\n')

	fmt.Println(verify(message,magazine))

}


