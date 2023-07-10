package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
		text := readFile(input)
		text = correct(text)
		file.WriteString(text)
		defer file.Close()
	}
}
func correct(s string) string {
	n := SplitWhiteSpaces(s)
	search(n, "(hex)", hex)
	search(n, "(bin)", bin)
	search(n, "(low)", low)
	for x := 1; x <= 10; x++ {
		search2(n, "(low,", x, x, low)
	}
	search(n, "(up)", up)
	for x := 1; x <= 10; x++ {
		search2(n, "(up,", x, x, up)
	}
	search(n, "(cap)", cap)
	for x := 1; x <= 10; x++ {
		search2(n, "cap,", x, x, cap)
	}
	z := strings.Join(n, " ")
	z = strings.ReplaceAll(z, " (hex)", "")
	z = strings.ReplaceAll(z, " (bin)", "")
	z = strings.ReplaceAll(z, " (low)", "")
	z = strings.ReplaceAll(z, " hex)", "")
	z = strings.ReplaceAll(z, " (cap)", "")
	z = strings.ReplaceAll(z, " (up)", "")
	for x := 1; x <= 10; x++ {
		z = replace(z, "(cap, ", x)
		z = replace(z, "(up, ", x)
		z = replace(z, "(low, ", x)
	}
	z = strings.ReplaceAll(z, ",", ", ")
	x := SplitWhiteSpaces(z)
	x = dothe(x)
	x = format(x, ",")
	x = format(x, ",")
	x = format(x, "?")
	x = format(x, "!")
	x = format(x, "…")
	x = format(x, ":")
	x = format(x, ";")
	x = format(x, ".")
	x = ancorrect(x)
	b := strings.Join(x, " ")
	return b
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
	z, _ := strconv.ParseUint(text, 16, 64)
	return strconv.Itoa(int(z))
}

func bin(text string) string {
	z, _ := strconv.ParseUint(text, 2, 64)
	return strconv.Itoa(int(z))
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
	s := strings.Title(text)
	return s

}

func an(text string) bool {
	if text[0] == 'a' {
		return true
	}
	return false
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

func SplitWhiteSpaces(s string) []string {
	var res []string
	var x string
	snew := s + " "
	for _, c := range snew {
		if c == ' ' || c == '	' || c == '\n' {
			if x != "" {
				res = append(res, x)
				x = ""
			}
		} else {
			x += string(c)
		}
	}
	return res
}

func search(n []string, sep string, function func(s string) string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(n[x], sep) {
			n[x-1] = function(n[x-1])
		}
	}
	return n
}

func search2(n []string, sep string, sepp, v int, function func(s string) string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(n[x], sep) {
			num := strconv.Itoa(sepp)
			if x+1 < len(n) {
				if strings.Contains(n[x+1], num) {
					for z := 1; z <= v; z++ {
						if z <= x {
							n[x-z] = function(n[x-z])
						}
					}
				}
			}
		}
	}
	return n
}

func replace(z, n string, x int) string {

	num := strconv.Itoa(x)
	n += num + ")"
	z = strings.ReplaceAll(z, n, "")
	return z
}
func format(n []string, a string) []string {
	var newArr []string
	for x := 0; x < len(n); x++ {
		if x != len(n)-1 {
			if strings.Contains(string(n[x+1][0]), a) {
				b := (n[x] + n[x+1])
				newArr = append(newArr, b)
				x++
			} else {
				newArr = append(newArr, n[x])
			}
		} else {
			newArr = append(newArr, n[x])
		}
	}
	return newArr
}

func ancorrect(n []string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(string(n[x][0]), "a") || strings.Contains(string(n[x][0]), "u") || strings.Contains(string(n[x][0]), "i") || strings.Contains(string(n[x][0]), "e") || strings.Contains(string(n[x][0]), "o") {
			if an(n[x-1]) && len(n[x-1]) == 1 {
				n[x-1] += "n"
			}
		}
	}
	return n
}

func dothe(n []string) []string {
	var newArr []string
	for x := 0; x < len(n); x++ {
		newArr = append(newArr, n[x])
		if strings.Contains(n[x], ":") && n[x+1] == "‘" {
			b := (n[x+1] + n[x+2])
			newArr = append(newArr, b)
			x++
			x++
		} else if x+1 < len(n) && n[x+1] == "‘" {
			b := (n[x] + n[x+1])
			newArr[x-1] = b
			x++
		}
	}
	return newArr
}
