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
	fmt.Print(len(n))
	n = ancorrect(n)
	for x := 1; x <= 4000; x++ {
		search2(n, "(low,", x, x, low)
	}
	search(n, "(up)", up)
	for x := 1; x <= 4000; x++ {
		search2(n, "(up,", x, x, up)
	}
	search(n, "(cap)", cap)
	for x := 1; x <= 4000; x++ {
		search2(n, "cap,", x, x, cap)
	}
	search(n, "(hex)", hexdecimal)
	search(n, "(bin)", Binary)
	search(n, "(low)", low)
	n = cotations('\'', n)
	n = cotations('"', n)
	z := strings.Join(n, " ")
	z = strings.ReplaceAll(z, "(hex)", "")
	z = strings.ReplaceAll(z, "(bin)", "")
	z = strings.ReplaceAll(z, "(low)", "")
	z = strings.ReplaceAll(z, "(hex)", "")
	z = strings.ReplaceAll(z, "(cap)", "")
	z = strings.ReplaceAll(z, "(up)", "")
	z = strings.ReplaceAll(z, " .", ".")
	for x := 1; x <= 4000; x++ {
		z = replace(z, "(cap, ", x)
		z = replace(z, "(up, ", x)
		z = replace(z, "(low, ", x)
	}
	z = strings.ReplaceAll(z, ",", " , ")
	x := strings.Split(z, " ")
	x = deleteNil(x)
	x = format2(x, "(hex)")
	x = format2(x, "(bin)")
	x = format2(x, "(cap)")
	b := strings.Join(x, " ")
	d := strings.Split(b, " ")
	h := strings.Join(d, " ")
		h = strings.ReplaceAll(h," ,",",")


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

func up(text string) string {
	return strings.ToUpper(text)
}

func low(text string) string {
	return strings.ToLower(text)
}

func cap(text string) string {
	x := low(text)
	for i := 0; i < len(text); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			strings.ToUpper(string(rune(x[i])))
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

func ancorrect(n []string) []string {
	for x := 1 ; x < len(n); x++ {
		if strings.Contains(strings.ToLower(string(n[x][0])), "o") || strings.Contains(strings.ToLower(string(n[x][0])), "u")  || strings.Contains(strings.ToLower(string(n[x][0])), "i")  || strings.Contains(strings.ToLower(string(n[x][0])), "e")  || strings.Contains(strings.ToLower(string(n[x][0])), "a")    {
			if an(n[x-1]) && len(n[x-1]) == 1 {
				n[x-1] += "n"
			}
		}
	}
	return n
}

func format2(n []string, a string) []string {
	var newArr []string
	for x := 0; x < len(n); x++ {
		if strings.Contains(n[x], a) {
			x++
		} else {
			newArr = append(newArr, n[x])
		}
	}
	return newArr
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
			fmt.Print(cond)
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