package main

import (
	"math"
)

var (
	ones     = []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	onesF    = []string{"", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
	teens    = []string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens     = []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds = []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
)

// numToStr возвращает сумму прописью с рублями и копейками
func numToStr(amount float64) string {
	rub := int(amount)
	kop := int(math.Round((amount - float64(rub)) * 100))

	parts := []string{}

	// Обрабатываем тысячи отдельно
	if rub >= 1000 {
		th := rub / 1000
		if th > 0 {
			parts = append(parts, tripletToStr(th, true))
			parts = append(parts, thousandSuffix(th))
		}
		rub %= 1000
	}

	if rub > 0 {
		parts = append(parts, tripletToStr(rub, false))
		parts = append(parts, rubleSuffix(rub))
	} else if len(parts) == 0 {
		parts = append(parts, "ноль рублей")
	} else {
		parts = append(parts, rubleSuffix(rub))
	}

	// Копейки тоже прописью (женский род для единиц)
	parts = append(parts, tripletToStr(kop, true))
	parts = append(parts, kopSuffix(kop))

	final := ""
	for _, p := range parts {
		if p != "" {
			final += p + " "
		}
	}
	return final[:len(final)-1]
}

func tripletToStr(n int, female bool) string {
	if n == 0 {
		return "ноль"
	}

	h := n / 100
	t := (n % 100) / 10
	o := n % 10

	words := []string{}
	if h > 0 {
		words = append(words, hundreds[h])
	}

	if t == 1 {
		words = append(words, teens[o])
	} else {
		if t > 1 {
			words = append(words, tens[t])
		}
		if o > 0 {
			if female {
				words = append(words, onesF[o])
			} else {
				words = append(words, ones[o])
			}
		}
	}

	return joinNonEmpty(words)
}

func thousandSuffix(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return "тысяча"
	} else if n%10 >= 2 && n%10 <= 4 && !(n%100 >= 12 && n%100 <= 14) {
		return "тысячи"
	} else {
		return "тысяч"
	}
}

func rubleSuffix(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return "рубль"
	} else if n%10 >= 2 && n%10 <= 4 && !(n%100 >= 12 && n%100 <= 14) {
		return "рубля"
	} else {
		return "рублей"
	}
}

func kopSuffix(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return "копейка"
	} else if n%10 >= 2 && n%10 <= 4 && !(n%100 >= 12 && n%100 <= 14) {
		return "копейки"
	} else {
		return "копеек"
	}
}

func joinNonEmpty(parts []string) string {
	res := ""
	for _, p := range parts {
		if p != "" {
			res += p + " "
		}
	}
	return res[:len(res)-1]
}
