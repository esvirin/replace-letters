package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Для выхода введите exit")
	fmt.Println("Введите текст для замены символов:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		fmt.Println("Результат: \n" + replace(input) + "\n")
	}
}

func replace(input string) string {
	replacement := map[rune]string{
		'а': "a",
		'о': "o",
		'у': "y",
		'е': "e",
		'с': "c",
		'т': "t",
		'a': "а",
		'o': "о",
		'y': "у",
		'e': "е",
		'c': "с",
		't': "т",
	}

	result := strings.Builder{}

	for _, letter := range input {
		if rep, ok := replacement[letter]; ok {
			result.WriteString(rep)
			continue
		}

		result.WriteRune(letter)
	}
	return result.String()
}
