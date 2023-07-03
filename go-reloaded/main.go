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
	search(n, "a", an)
	for x := 1; x <= 10; x++ {
		search2(n, "(low,", string(x)+")", x, low)
	}
	search(n, "(up)", up)
	for x := 1; x <= 10; x++ {
		search2(n, "(up,", string(x)+")", x, up)
	}
	search(n, "(cap)", cap)
	for x := 1; x <= 10; x++ {
		search2(n, "cap," , string(x) , x , cap)
	}
	z := strings.Join(n, " ")
	z = strings.ReplaceAll(z, " (hex)", "")
	z = strings.ReplaceAll(z, " (bin)", "")
	z = strings.ReplaceAll(z, " (low)", "")
	z = strings.ReplaceAll(z, " hex)", "")
	z = strings.ReplaceAll(z, " (cap)", "")
	z = strings.ReplaceAll(z, " (up)", "")
	z = strings.ReplaceAll(z, "A ", "")
	return z

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

func an(text string) string {
	if text[0] == 'a' {
		text = "an " + text
	}
	return text
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

func search2(n []string, sep , sepp string, v int, function func(s string) string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(n[x], sep) {
			if strings.Contains(n[x+1], sepp) {
				for z := 1; z <= v ; z++ {
					n[x-z] = function(n[x-z]) 
				}
			}
		}
	}
	return n
}

