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
	sections := []string{""}
	totals := []float64{0}

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// Extract the unit we're dealing in
		if line == 0 && strings.Contains(text, "# UNIT:") {
			unit = strings.TrimSpace(strings.TrimPrefix(text, "# UNIT:"))
			continue
		}

		// Skip comments, but build sections
		if strings.HasPrefix(text, "#") {
			if !strings.HasSuffix(text, ":") {
				continue
			}

			section := strings.Trim(text, " #:")

			sections = append(sections, section)
			totals = append(totals, 0)
		}

		// Strip the comment away from the entry
		if strings.Contains(text, "#") {
			text = strings.Split(text, "#")[0]
		}

		// Extract the number
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		submatchall := re.FindAllString(text, -1)
		for _, element := range submatchall {
			val, _ := strconv.ParseFloat(element, 64)
			totals[len(totals)-1] += val
		}
		line += 1
	}

	p := message.NewPrinter(language.English)

	for i, section := range sections {
		// Don't print empty sections or ones that are near zero
		if section == "" || (totals[i] > -0.001 && totals[i] < 0.001) {
			continue
		}

		p.Printf("%s: %d %s\n", section, number.Decimal(totals[i]), unit)
	}

	total := 0.0
	for _, n := range totals {
		total += n
	}
	if len(sections) > 1 {
		p.Printf("=====\n")
	}
	p.Printf("TOTAL: %d %s\n", number.Decimal(total), unit)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
