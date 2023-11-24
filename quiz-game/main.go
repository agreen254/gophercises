package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fname := flag.String("f", "problems.csv", "optional custom file name to read problems from")
	timeLimit := flag.Int("t", 30, "specify the time limit")
	// shouldShuffle := flag.Bool("s", false, "specify if you want the problems shuffled or not")
	flag.Parse()

	file, err := processFileInput(*fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	handleStart(*timeLimit)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	inReader := bufio.NewReader(os.Stdin)
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Could not process records.", err)
	}

	var numCorrect = 0
	go func() {
		<-timer.C
		fmt.Printf("\nelapsed time finished!\nyou got %d/%d correct\n", numCorrect, len(records))
		os.Exit(0)
	}()

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
	fmt.Printf("You got %d/%d correct. ", numCorrect, len(records))
	if numCorrect == len(records) {
		fmt.Printf("Nice job!\n")
	} else if numCorrect == 0 {
		fmt.Printf("Ouch.\n")
	} else {
		fmt.Printf("\n")
	}

	os.Exit(0)
}

func handleStart(limit int) {
	fmt.Printf("press Enter to begin quiz (limit %d seconds): ", limit)
	beginReader := bufio.NewReaderSize(os.Stdin, 1)
	for true {
		input, err := beginReader.ReadByte()
		if err != nil {
			log.Fatalln(err)
		}
		// carriage return || line break
		if input == 13 || input == 10 {
			break
		}
	}
}

func processFileInput(fname string) (file *os.File, err error) {
	return os.Open(fname)
}
