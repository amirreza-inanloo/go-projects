package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "CSV file in 'question,answer' format")
	timeLimit := flag.Int("limit", 10, "Time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "Shuffle the quiz questions")
	flag.Parse()

	// Open the CSV file
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Failed to open the CSV file:", err)
		return
	}
	defer file.Close()

	// Read all records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed to read the CSV file:", err)
		return
	}

	// Shuffle questions if flag is true
	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	// Prompt user to start
	fmt.Printf("Press Enter to start the quiz (you have %d seconds)...", *timeLimit)
	bufio.NewReader(os.Stdin).ReadString('\n')

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemloop:
	for i, record := range records {
		if len(record) != 2 {
			continue
		}
		question := strings.TrimSpace(record[0])
		answer := strings.TrimSpace(strings.ToLower(record[1]))

		fmt.Printf("Question %d: %s\n", i+1, question)
		answerCh := make(chan string)

		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := strings.TrimSpace(strings.ToLower(scanner.Text()))
			answerCh <- text
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			break problemloop
		case userAnswer := <-answerCh:
			if userAnswer == answer {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(records))
}
