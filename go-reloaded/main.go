package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	osNo := len(os.Args)
	if osNo != 3 {
		return
	} else {
		input := os.Args(1)
		output := os.Args(2)
		file, err := os.Create(output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer file.close()
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
		continue
	}
	f, err := ioutil.ReadAll(file)
	if err != nil {
		printError(err)
		continue
	}
	text = f
	file.Close()
	return text
}

func hex(text string) string {
	var z int
	for _, c := range text {
		power := 0
		if c >= '0' && c <= '9' {
			x := piscine.Atoi(string(c))
			z += c * IterativePower(16, power)
			power++
		} else {
			x := piscine.Atoi(string((c - 'A')))
			z += c * IterativePower(16, power)
			power++
		}
	}
	return strconv.Itoa(z)
}

func bin(text string) string {


}

func up(text string) string {

}

func low(text string) string {

}

func cap(text string) string {

}

func an(text string) string {

}

func punctuations(text string) string {

}
