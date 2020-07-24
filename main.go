package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Error")
		return
	}

	fs := ""
	flag := "jdjwqhdwqhdkqwklwqdhkwqdhwkqdh"
	arg := os.Args[1]
	for _, v := range arg {
		if !(v >= 32 && v <= 126) {
			return
		}
	}

	if len(args) == 2 {
		fs = os.Args[2]
		if fs != "standard" && fs != "shadow" && fs != "thinkertoy" {
			fmt.Println("Wrong Font!")
			return
		}
	}

	if len(args) == 3 {
		fs = os.Args[2]
		if fs != "standard" && fs != "shadow" && fs != "thinkertoy" {
			fmt.Println("Wrong Font!")
			return
		}
		flag = os.Args[3]

		if len(flag) < 14 {
			fmt.Println("Wrong command or type of format!")
			return
		}
		if flag[0:9] != "--output=" || !(flag[len(flag)-4:len(flag)] == ".txt") {
			fmt.Println("Wrong command or type of format!")
			return
		}
		CreateFile(flag[9:len(flag)])
	}
	formattype := FormatType(fs)
	file, err := os.Open(formattype)

	if len(args) > 3 {
		fmt.Println("Too many arguments!")
		return
	}

	// formattype := FormatType(fs)
	// file, err := os.Open(formattype)

	if err != nil {
		fmt.Println("Error")
		return
	}
	defer file.Close()
	banners := [][]string{}
	banner := []string{}

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		i++
		banner = append(banner, scanner.Text())

		if i == 9 {
			banners = append(banners, banner)
			banner = []string{}
			i = 0
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error")
		return
	}
	// CreateFile(flag[9:len(flag)])

	array := strings.Split(arg, "\\n")
	for i := 0; i < len(array); i++ {
		for j := 1; j <= 8; j++ {
			for _, value := range array[i] {

				str := banners[int(value)-32]
				PrintOutput(str[j], flag[9:len(flag)])
				fmt.Print(str[j])
			}
			if len(array[i]) != 0 {
				PrintOutput("\n", flag[9:len(flag)])
				fmt.Println()
			}
		}
	}

}

func FormatType(fs string) string {
	if fs == "shadow" {
		return "shadow.txt"
	}
	if fs == "thinkertoy" {
		return "thinkertoy.txt"
	}

	return "standard.txt"
}
func CreateFile(str string) {
	file, err := os.Create(str)
	if err != nil {
	}
	defer file.Close()
}

func PrintOutput(str string, flag string) {

	file, err := os.OpenFile(flag, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {

		return

	}
	file.WriteString(str)
	defer file.Close()
}
