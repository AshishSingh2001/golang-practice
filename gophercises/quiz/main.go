package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func readCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file "+filePath, err)
	}

	csvReader := csv.NewReader(f)
	questions, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV for "+filePath, err)
	}
	return questions
}

func main() {
	filePath := flag.String("file", "problems.csv", "file location for questions")
	randomise := flag.Bool("randomise", false, "Randomise input deafaults to false")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz")
	flag.Parse()

	// read parse and randomise
	r := readCsv(*filePath)
	ques := parseProblems(r)
	rand.Seed(time.Now().UnixNano())
	if *randomise {
		rand.Shuffle(len(ques), func(i, j int) {
			ques[i], ques[j] = ques[j], ques[i]
		})
	}

	// ask and score
	var score int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, v := range ques {
		// make a channel and take the input through it
		// we want to stop the input so we would like
		// it to not block the main thread
		answerCh := make(chan string)
		go func() {
			fmt.Print("Question", i, "\n", v.q, " = ")
			var inp string
			fmt.Scan(&inp)
			inp = strings.ToLower(strings.Trim(inp, " "))
			answerCh <- inp
		}()
		select {
		case <-timer.C:
			fmt.Println("\nTime's Up!!!!!!!!!")
			fmt.Println("Your Score is : ", score, "out of", len(ques))
			return
		case inp := <-answerCh:
			if inp == v.a {
				score++
			}
		}
	}
	fmt.Println("Your Score is : ", score, "out of", len(ques))
}

func parseProblems(r [][]string) []Problem {
	ret := make([]Problem, len(r))
	for i, v := range r {
		ret[i] = Problem{q: v[0], a: strings.ToLower(strings.TrimSpace(v[1]))}
	}
	return ret
}

type Problem struct {
	q string
	a string
}
