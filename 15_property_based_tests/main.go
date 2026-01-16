package main

import "strings"

func main() {
	// entry point
	ConvertToRoman(39)
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	/*
		switch true {
		case (number >= 1 && number <= 3):
			for range number {
				result.WriteString("I")
			}
			return result.String()
		// return  strings.Repeat("I", number)

		case (number == 4):
			return "IV"

		case (number >= 5 && number <= 8):
			countGreaterThan5 := number - 5
			if countGreaterThan5 == 0 {
				return "V"
			} else {
				result.WriteString("V")
				for range countGreaterThan5 {
					result.WriteString("I")
				}
				return result.String()
			}
		case number > 8:
			result.WriteString("IX")
			return result.String()
		default:
			return ""
		}
	*/

	// classic greedy algorithm
	// Always subtract the largest possible Roman value until you can’t anymore, then move to the next smaller one.
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var arabic = 0
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
