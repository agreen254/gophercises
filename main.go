package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fname := flag.String("f", "problems.csv", "optional custom file name to read problems from")
	flag.Parse()

	file, err := processFileInput(*fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	inReader := bufio.NewReader(os.Stdin)
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Could not process records.", err)
	}

	var numCorrect = 0
	for _, record := range records {
		fmt.Print(record[0] + " = ")
		input, err := inReader.ReadString('\n')
		if err != nil {
			log.Fatalln("An error occurred while reading input.", err)
		}
		if strings.TrimSpace(input) == record[1] {
			numCorrect++
		}
	}

	fmt.Printf("You got %d/%d correct.\n", numCorrect, len(records))

}

func processFileInput(fname string) (file *os.File, err error) {
	return os.Open(fname)
}
