package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	osNo := len(os.Args)
	if osNo != 3 {
		fmt.Println("error , os.Args != 3")
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
	s = strings.ReplaceAll(s, "' ", " ' ")
	s = strings.ReplaceAll(s, " '", " ' ")
	s = strings.ReplaceAll(s, "’", " ’ ")
	s = strings.ReplaceAll(s, "?", " ? ")
	s = strings.ReplaceAll(s, "!", " ! ")
	s = strings.ReplaceAll(s, ":", " : ")
	s = strings.ReplaceAll(s, "\"", " \" ")
	s = strings.ReplaceAll(s, ";", " ; ")
	s = strings.ReplaceAll(s, ".", " . ")

	n := strings.Split(s, " ")
	n = deleteNil(n)
	n = ancorrect(n)

	search22(n, "(low,", strings.ToLower)
	search(n, "(up)", strings.ToUpper)
	search22(n, "(up,", strings.ToUpper)
	search(n, "(cap)", cap)
	search22(n, "cap,", cap)

	search(n, "(hex)", hexdecimal)
	search(n, "(bin)", Binary)
	search(n, "(low)", strings.ToLower)

	n = cotations('\'', n)
	n = cotations('"', n)
	n = cotations('’', n)

	z := strings.Join(n, " ")
	z = strings.ReplaceAll(z, "(hex)", "")
	z = strings.ReplaceAll(z, "(bin)", "")
	z = strings.ReplaceAll(z, "(low)", "")
	z = strings.ReplaceAll(z, "(hex)", "")
	z = strings.ReplaceAll(z, "(cap)", "")
	z = strings.ReplaceAll(z, "(up)", "")
	z = strings.ReplaceAll(z, " .", ".")
	z = strings.ReplaceAll(z, ",", " , ")
	x := strings.Split(z, " ")
	x = deleteNil(x)
	b := strings.Join(x, " ")
	d := strings.Split(b, " ")
	h := strings.Join(d, " ")
	h = strings.ReplaceAll(h, " ,", ",")
	s = strings.ReplaceAll(s, " ?", "?")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " :", ":")
	s = strings.ReplaceAll(s, " \"", "\"")
	s = strings.ReplaceAll(s, " ;", ";")
	s = strings.ReplaceAll(s, " .", ".")

	return h
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

func hexdecimal(text string) string {
	z, err := strconv.ParseUint(text, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(z))

}

func Binary(text string) string {
	z, err := strconv.ParseUint(text, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(int(z))

}

func cap(text string) string {
	x := strings.ToLower(text)
	for i := 0; i < len(text); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			x = x[:i] + strings.ToUpper(string(x[i])) + x[i+1:]
			return x
		}
	}
	return x
}

func an(text string) bool {
	if text[0] == 'a' || text[0] == 'A' {
		return true
	}
	return false
}
func printError(err error) {
	fmt.Println("ERROR: " + err.Error())
	os.Exit(1)
}

func search(n []string, sep string, function func(s string) string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(n[x], sep) {
			if x-1 >= 0 {
				n[x-1] = function(n[x-1])
			}
		}
	}
	return n
}

func search22(n []string, sep string, function func(s string) string) []string {
	for x := 0; x < len(n); x++ {
		if strings.Contains(n[x], sep) {
			n[x] = ""
			z := strings.TrimRight(n[x+1], ")")
			n[x+1] = ""
			num, err := strconv.Atoi(z)
			if err != nil {
				printError(err)
			}
			for z := 1; z <= num; z++ {
				if z <= x {
					n[x-z] = function(n[x-z])
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

func ancorrect(n []string) []string {
	for x := 1; x < len(n); x++ {
		if strings.Contains(strings.ToLower(string(n[x][0])), "o") || strings.Contains(strings.ToLower(string(n[x][0])), "u") || strings.Contains(strings.ToLower(string(n[x][0])), "i") || strings.Contains(strings.ToLower(string(n[x][0])), "e") || strings.Contains(strings.ToLower(string(n[x][0])), "a") {
			if an(n[x-1]) && len(n[x-1]) == 1 {
				n[x-1] += "n"
			}
		}
	}
	return n
}

func cotations(n rune, s []string) []string {
	cond := false
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 0 && rune(s[i][0]) == n && !cond && i+1 < len(s) {
			if len(s[i+1]) > 0 && (rune(s[i+1][0]) == n) {
				cond = false
			} else {
				cond = true
			}
			s[i+1] = string(n) + s[i+1]
			s[i] = ""
			i++
		} else if len(s[i]) > 0 && rune(s[i][0]) == n && cond {
			s[i-1] = s[i-1] + string(n)
			s[i] = ""
			cond = false
		}
	}
	return s
}

func deleteNil(s []string) []string {
	var newArr []string
	for x := 0; x < len(s); x++ {
		if s[x] != "" {
			newArr = append(newArr, s[x])
		}
	}
	return newArr
}
