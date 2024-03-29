package main

// Задание 26
// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//
// 	Например:
// 		abcd — true
//		abCdefAaf — false
// 		aabcd — false

import (
	"fmt"
	"strings"
)

func uniqueChars(s string) bool {
	hasRepeat := make(map[rune]bool)
	for _, r := range strings.ToLower(s) {

		if hasRepeat[r] {
			return false
		}
		hasRepeat[r] = true
	}
	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aAbcd", "Эй,жлоб!Гдетуз?Прячь юныхсъёмщицвшкаф."}
	for _, s := range arr {
		fmt.Printf("%s \t%t\n", s, uniqueChars(s))
	}
}
