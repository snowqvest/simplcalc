package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	oper := ""
	switch {
	case strings.Contains(input, "+"):
		oper = "+"
	case strings.Contains(input, "-"):
		oper = "-"
	case strings.Contains(input, "*"):
		oper = "*"
	case strings.Contains(input, "/"):
		oper = "/"
	default:
		fmt.Println("Error: Wrong expression format (Example: 3+2)")
		return
	}

	nums := strings.Split(input, oper)

	if len(nums) != 2 {
		fmt.Println("Error: Wrong expression format (Example: 3+2)")
		return
	}

	if strings.ContainsAny(nums[0], "+-*/") || strings.ContainsAny(nums[1], "+-*/") {
		fmt.Println("Error: Wrong expression format (Example: 3+2)")
		return
	}

	nums[0] = strings.TrimSpace(nums[0])
	nums[1] = strings.TrimSpace(nums[1])

	num1 := romanToInt(nums[0])
	num2 := romanToInt(nums[1])
	isRoman := true

	if num1 == 0 || num2 == 0 {
		var err1, err2 error
		isRoman = false
		num1, err1 = strconv.Atoi(nums[0])
		num2, err2 = strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error: Only Roman or decimal numerals allowed. Not both types in single expression. Both numbers must be in range from 1 to 10.")
			return
		}
	}

	if num1 > 10 || num2 > 10 || num1 < 1 || num2 < 1 {
		fmt.Println("Error: Both numbers must be in range from 1 to 10.")
		return
	}

	res := 0
	switch oper {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2

	}

	if isRoman {
		if res <= 0 {
			fmt.Println("Error: Negative numbers (or zero) can't be written in Roman numeral system.")
			return
		}
		fmt.Println(intToRoman(res))
	} else {
		fmt.Println(res)
	}

}

func romanToInt(str string) int {
	romans := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	res := 0
	for i := 0; i < 10; i++ {
		if str == romans[i] {
			res = i + 1
		}
	}
	return res

}

func intToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
