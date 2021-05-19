package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a file path.")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// If a unit is specified, read it
	unit := ""
	total := 0.0

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		if line == 0 && strings.Contains(text, "# UNIT: ") {
			unit = strings.TrimPrefix(text, "# UNIT: ")
			continue
		}

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		submatchall := re.FindAllString(text, -1)
		for _, element := range submatchall {
			val, _ := strconv.ParseFloat(element, 64)
			total += val
		}
		line += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	p := message.NewPrinter(language.English)
	p.Printf("TOTAL: %d %s\n", number.Decimal(total), unit)
}
