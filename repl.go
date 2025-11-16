package main

import "strings"

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	splitStrings := strings.Split(text, " ")
	for index, str := range splitStrings {
		splitStrings[index] = strings.ToLower(strings.Trim(str, " "))
	}
	return splitStrings
}
