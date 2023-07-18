package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman, arabic bool

func recognizenum(in string) (int, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		if arabic == true {
			return 0, errors.New("YOU CAN USE ARABIC OR ROMAN NOT BOTH")
		}
		roman = true
		switch in {
		case "I":
			return 1, nil
		case "II":
			return 2, nil
		case "III":
			return 3, nil
		case "IV":
			return 4, nil
		case "V":
			return 5, nil
		case "VI":
			return 6, nil
		case "VII":
			return 7, nil
		case "VIII":
			return 8, nil
		case "IX":
			return 9, nil
		case "X":
			return 10, nil
		default:
			return 0, errors.New("UNKNOWN NUMBER")
		}
	}
	if out > 10 || out < 1 {
		return 0, errors.New("input numbers from 1 to 10 only handled")
	}
	if roman == true {
		return 0, errors.New("YOU CAN USE ARABIC OR ROMAN NOT BOTH")
	}
	arabic = true
	return out, nil
}

func arabictoroman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid number"
	}

	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(arabic); i++ {
		for num >= arabic[i] {
			result += roman[i]
			num -= arabic[i]
		}
	}

	return result
}

func calculate(operand string, a, b int) (string, error) {
	var result int
	switch operand {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("cannot divide by zero")
		}
		result = a / b
	default:
		return "", errors.New("UNKNOWN OPERATION")
	}
	if roman == true && result < 1 {
		return "", errors.New("ROMAN NOTATRION NOT SUPPORT NEGATIVE NUMBERS")
	}
	if roman == true {
		return arabictoroman(result), nil
	}
	return strconv.Itoa(result), nil

}

func main() {
	fmt.Println("type your expression")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	line = strings.TrimSpace(line)
	arr := strings.Split(line, " ")
	//fmt.Println(arr)
	if 3 != len(arr) {
		fmt.Println("INVALID OPERATION FORMAT")
		os.Exit(1)
	}
	a := arr[0]
	b := arr[2]
	sign := arr[1]
	//fmt.Println(a, b, sign)

	aNum, err := recognizenum(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bNum, err := recognizenum(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := calculate(sign, aNum, bNum)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
