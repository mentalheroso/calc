package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numberString1, sign, numberString2, sum string
	var numberRome1Int, numberRome2Int, res int
	romeToArab := map[string]int{
		"I":      1,
		"II":     2,
		"III":    3,
		"IV":     4,
		"V":      5,
		"VI":     6,
		"VII":    7,
		"VIII":   8,
		"IX":     9,
		"X":      10,
		"XI":     11,
		"XII":    12,
		"XIII":   13,
		"XIV":    14,
		"XV":     15,
		"XVI":    16,
		"XVII":   17,
		"XVIII":  18,
		"XIX":    19,
		"XX":     20,
		"XXI":    21,
		"XXIV":   24,
		"XXV":    25,
		"XXVII":  27,
		"XXVIII": 28,
		"XXX":    30,
		"XXXII":  32,
		"XXXV":   35,
		"XXXVI":  36,
		"XL":     40,
		"XLII":   42,
		"XLV":    45,
		"XLVIII": 48,
		"XLIX":   49,
		"L":      50,
		"LIV":    54,
		"LVI":    56,
		"LX":     60,
		"LXIII":  63,
		"LXIV":   64,
		"LXX":    70,
		"LXXII":  72,
		"LXXX":   80,
		"LXXXI":  81,
		"XC":     90,
		"C":      100,
	}
	arabToRome := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		24:  "XXIV",
		25:  "XXV",
		27:  "XXVII",
		28:  "XXVIII",
		30:  "XXX",
		32:  "XXXII",
		35:  "XXXV",
		36:  "XXXVI",
		40:  "XL",
		42:  "XLII",
		45:  "XLV",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		54:  "LIV",
		56:  "LVI",
		60:  "LX",
		63:  "LXIII",
		64:  "LXIV",
		70:  "LXX",
		72:  "LXXII",
		80:  "LXXX",
		81:  "LXXXI",
		90:  "XC",
		100: "C",
	}
	fmt.Println("Введите выражение")
	fmt.Scan(&numberString1, &sign, &numberString2)
	sum = numberString1 + numberString2
	math123 := true
	if math123 == strings.ContainsAny(sum, "/*+-") {
		panic("строка не является математической операцией")
	}
	if strings.ContainsAny(sum, "IVX") && strings.ContainsAny(sum, "1234567890") {
		panic("используются одновременно разные системы счисления")
	}
	i, j := 0, 0
	if strings.ContainsAny(sum, "IVX") {
		for {
			i++
			if numberString1 == arabToRome[i] {
				numberRome1Int = romeToArab[numberString1]
				break
			}
		}
		for {
			j++
			if numberString2 == arabToRome[j] {
				numberRome2Int = romeToArab[numberString2]
				break
			}
		}
		raz := numberRome1Int - numberRome2Int
		if 0 > numberRome1Int || numberRome1Int > 10 || 0 > numberRome2Int || numberRome2Int > 10 {
			panic("числа вне диапозона[I:X]")
		}
		switch sign {
		case "+":
			res = numberRome1Int + numberRome2Int
		case "-":
			if raz < 0 || raz == 0 {
				panic("Разность равна или меньше 0 (в римской системе нет отрицательных чисел и 0)")
			}
			res = numberRome1Int - numberRome2Int
		case "*":
			res = numberRome1Int * numberRome2Int
		case "/":
			res = numberRome1Int / numberRome2Int
		default:
			panic("Нет знака")
		}
		for key, value := range romeToArab {
			if value == res {
				fmt.Println(key)
			}
		}
	} else {
		number1Int, err1 := strconv.Atoi(numberString1)
		if err1 != nil {
			panic("Введены не числа")
		}
		number2Int, err1 := strconv.Atoi(numberString2)
		if err1 != nil {
			panic("Введены не числа")
		}
		if 1 <= (number1Int) && (number1Int) <= 10 && 1 <= (number2Int) && (number2Int) <= 10 {
			switch sign {
			case "+":
				fmt.Print(number1Int + number2Int)
			case "-":
				fmt.Print(number1Int - number2Int)
			case "*":
				fmt.Print(number1Int * number2Int)
			case "/":
			default:
				panic("Нет знака")
			}
		} else {
			panic("числа вне диапозона[1:10]")
		}
	}
}
