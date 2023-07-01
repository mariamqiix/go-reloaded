package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	osNo := len(os.Args)
	if osNo == 0 {
		return
	} else {
		input := os.Args[1]
		output := os.Args[2]
		file, err := os.Create(output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer file.Close()
		text := readFile(input)
		text = correct(text)

	}
}

func correct(text string) string {

}

func readFile(name string) string {
	var text string
	file, err := os.Open(name)
	if err != nil {
		printError(err)
	}
	f, err := ioutil.ReadAll(file)
	if err != nil {
		printError(err)
	}
	text = string(f)
	file.Close()
	return text
}

func hex(text string) string {
	var z int
	for _, c := range text {
		power := 0
		if c >= '0' && c <= '9' {
			x := Atoi(string(c))
			z += x * IterativePower(16, power)
			power++
		} else {
			x := Atoi(string((c - 'A')))
			z += x * IterativePower(16, power)
			power++
		}
	}
	return strconv.Itoa(z)
}

func bin(text string) string {

}

func up(text string) string {
	X := []rune(text)
	c := ""
	for i := 0; i < len(X); i++ {
		if X[i] <= 122 && X[i] >= 97 {
			c = c + string(X[i]-32)
		} else {
			c = c + string(X[i])
		}
	}
	return c
}

func low(text string) string {
	X := []rune(text)
	c := ""
	for i := 0; i < len(X); i++ {
		if X[i] <= 90 && X[i] >= 65 {
			c = c + string(X[i]+32)
		} else {
			c = c + string(X[i])
		}
	}
	return c
}

func cap(text string) string {
	c := ""
	if rune(text) <= 122 && rune(text) >= 97 {
		c = c + string(rune(text)-32)
	} else {
		c = c + string(rune(text))
	}
	return c
}

func an(text string) string {

}

func punctuations(text string) string {

}

func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}

func IterativePower(nb int, power int) int {
	nn := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power > 0 {
		for i := 1; i <= power; i++ {
			nn *= nb
		}
	}
	return nn
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	result := 0
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}
	result *= sign
	return result
}
