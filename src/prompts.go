package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// promptRangeNumber prompts the user to select a range number
func PromptRangeNumber(totalRanges int) int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		fmt.Printf("Escolha a carteira (1 a %d): ", totalRanges)
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		rangeNumber, err := strconv.Atoi(input)
		if err == nil && rangeNumber >= 1 && rangeNumber <= totalRanges {
			return rangeNumber
		}
		fmt.Println("Numero invalido.")
	}
}

func PromptCPUNumber() int {
	reader := bufio.NewReader(os.Stdin)
	charReadline := '\n'

	if runtime.GOOS == "windows" {
		charReadline = '\r'
	}

	for {
		fmt.Printf("Quantos CPUs gostaria de usar?: ")
		input, _ := reader.ReadString(byte(charReadline))
		input = strings.TrimSpace(input)
		cpusNumber, err := strconv.Atoi(input)
		if err == nil && cpusNumber >= 1 && cpusNumber <= 50 {
			return cpusNumber
		}
		fmt.Println("Numero invalido.")
	}
}
