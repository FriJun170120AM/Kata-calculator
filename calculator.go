package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"bufio"
)

var roman, arabic bool

func recognizenum(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		if arabic == true {
			fmt.Println("YOU CAN USE ARABIC OR ROMAN NOT BOTH")
			os.Exit(1)
		}
		roman = true
		switch in {
		case "I":
			return 1
		case "II":
			return 2
		case "III":
			return 3
		case "IV":
			return 4
		case "V":
			return 5
		case "VI":
			return 6
		case "VII":
			return 7
		case "VIII":
			return 8
		case "IX":
			return 9
		case "X":
			return 10
		default:
			fmt.Println("UNKNOWN NUMBER")
			os.Exit(1)
        		// ... handle error
        		//panic(err)
		}
	}
	if out > 10 || out < 1 {
		fmt.Println("input numbers from 1 to 10 only handled")
		os.Exit(1)
	}
	if roman == true {
		fmt.Println("YOU CAN USE ARABIC OR ROMAN NOT BOTH")
		os.Exit(1)
	}
	arabic = true
	return out
}

func arabictoroman(arabicnumeral int) string {
	arabicmap := map[int]string{
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}
	result := ""
	for num, symbol := range arabicmap {
		for arabicnumeral >= num {
			result += symbol
			arabicnumeral -= num
		}
	}
	return result
}

func calculate(operand string, a, b int) string {
	var result int
	switch operand {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("UNKNOWN OPERATION")
		os.Exit(1)
	}
	if roman == true && result < 1 {
		fmt.Println("ROMAN NOTATRION NOT SUPPORT NEGATIVE NUMBERS")
		os.Exit(1)
	}
	if roman == true {
		return arabictoroman(result)	
	}
	return strconv.Itoa(result)
	
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
	a:= arr[0]
	b:= arr[2]
	sign:= arr[1]
	//fmt.Println(a, b, sign)

	fmt.Println(calculate(sign, recognizenum(a), recognizenum(b)))
}
