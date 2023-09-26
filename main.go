package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`^(.*?[*+-\/]{1})(.*)`)
var l = [20]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var arabic_numbers = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
var rome_numbers = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

var m map[string]string

func main() {
	// var first_digit, second_digit, operator int

	str := Scan1()
	if it_is_math_example(str) == true {
		str_before_operator := str_before_operator(str)
		str_after_operator := str_after_operator(str)
		operand := operand(str)
		if it_is_number_before_operator(str_before_operator, l) == true {
			if it_is_number_after_operator(str_after_operator, l) {
				if ns_is_same(str_before_operator, str_after_operator) {
					if it_is_arabic_number(str_before_operator, arabic_numbers) && it_is_arabic_number(str_after_operator, arabic_numbers) {
						fmt.Print(calc(str_before_operator, str_after_operator, operand))
					} else {
						fmt.Print(Roman(calc(str_before_operator, str_after_operator, operand)))
					}
				}
			}
		}
	}
}

func Scan1() string {
	fmt.Println("Введите пример:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return strings.ReplaceAll(in.Text(), " ", "")
}

func it_is_math_example(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "+" || string(s[i]) == "-" || string(s[i]) == "/" || string(s[i]) == "*" {
			count += 1
		}
	}
	if count < 1 {
		err := errors.New("Вывод ошибки, так как строка не является математической операцией")
		fmt.Print(err)
		return false
	} else if count > 1 {
		err := errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию")
		fmt.Print(err)
		return false
	}
	return true
}

func str_before_operator(s string) string {
	split := r.FindStringSubmatch(s)
	if len(split) < 2 {
		return ""
	}
	return (split[1][:len(split[1])-1])
}

func operand(s string) string {
	split := r.FindStringSubmatch(s)
	if len(split) < 2 {
		return ""
	}
	return (split[1][len(split[1])-1:])
}

func str_after_operator(s string) string {
	split := r.FindStringSubmatch(s)
	if len(split) < 2 {
		return ""
	}
	return split[2]
}

func it_is_number_before_operator(s string, list [20]string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	err := errors.New("Вывод ошибки, так как строка не является математической операцией")
	fmt.Print(err)
	return false
}

func it_is_number_after_operator(s string, list [20]string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	err := errors.New("Вывод ошибки, так как строка не является математической операцией")
	fmt.Print(err)
	return false
}

func ns_is_same(str_before_operator string, str_after_operator string) bool {
	if it_is_arabic_number(str_before_operator, arabic_numbers) == it_is_arabic_number(str_after_operator, arabic_numbers) {
		return true
	}
	err := errors.New("Вывод ошибки, так как используются одновременно разные системы счисления")
	fmt.Print(err)
	return false
}

func it_is_arabic_number(s string, list [10]string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func calc(str_before_operator string, str_after_operator string, operand string) int {
	calc := 0
	if it_is_arabic_number(str_before_operator, arabic_numbers) && it_is_arabic_number(str_after_operator, arabic_numbers) {
		number_one, _ := strconv.Atoi(str_before_operator)
		number_two, _ := strconv.Atoi(str_after_operator)
		if operand == "+" {
			calc := number_one + number_two
			return calc
		} else if operand == "-" {
			calc := number_one - number_two
			return calc
		} else if operand == "/" {
			calc := number_one / number_two
			return calc
		} else if operand == "*" {
			calc := number_one * number_two
			return calc
		}
	} else {
		number_one, _ := strconv.Atoi(rome_to_int(str_before_operator))
		number_two, _ := strconv.Atoi(rome_to_int(str_after_operator))
		if operand == "+" {
			calc := number_one + number_two
			return calc
		} else if operand == "-" {
			calc := number_one - number_two
			if calc <= 0 {
				err := errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и ")
				fmt.Print(err)
			}
			return calc
		} else if operand == "/" {
			calc := number_one / number_two
			if calc <= 0 {
				err := errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и ")
				fmt.Print(err)
			}
			return calc
		} else if operand == "*" {
			calc := number_one * number_two
			return calc
		}
	}
	return calc
}

func rome_to_int(num string) string {
	m = make(map[string]string)
	m["I"] = "1"
	m["II"] = "2"
	m["III"] = "3"
	m["IV"] = "4"
	m["V"] = "5"
	m["VI"] = "6"
	m["VII"] = "7"
	m["VIII"] = "8"
	m["IX"] = "9"
	m["X"] = "10"
	i := m[(num)]
	return i
}

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}
