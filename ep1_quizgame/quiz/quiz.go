package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	settings settings
	problems []Problem
	score    int
}

type settings struct {
	fileName string
	shuffle  bool
}

func MakeQuiz(fileName string, shuffle bool) *Quiz {
	return &Quiz{
		settings: settings{
			fileName: fileName,
			shuffle:  shuffle,
		},
		problems: make([]Problem, 0),
	}
}

func (q *Quiz) GetTotalQuestios() int {
	return len(q.problems)
}

func (q *Quiz) GetUserScore() int {
	return q.score
}

func (q *Quiz) ReadQuiz() {
	err := q.readFile(q.settings.fileName)
	if err != nil {
		log.Fatal("Error reading from file ", q.settings.fileName, " , err: ", err.Error())
	}
}

func (q *Quiz) BeginQuiz() {

	for i, problem := range q.problems {
		userAnswer, err := q.askQuestion(i, problem.question)
		if err != nil {
			log.Fatal("Error asking question, err: ", err.Error())
		}

		if userAnswer == problem.answer {
			q.score++
		}
	}
	fmt.Printf("Final Score %d/%d", q.score, len(q.problems))
}

func (q *Quiz) askQuestion(idx int, question string) (userAnswer string, err error) {

	buffReader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "Question %d # %s : ", idx+1, question)
	userAns, err := buffReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(userAns), nil
}

func (q *Quiz) readFile(filepath string) (err error) {

	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	csvFileReader := csv.NewReader(file)

	for {
		record, err := csvFileReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		q.problems = append(q.problems, Problem{
			question: record[0],
			answer:   record[1],
		})
	}

	if q.settings.shuffle {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for n := len(q.problems); n > 0; n-- {
			randIndex := r.Intn(n)

			q.problems[n-1], q.problems[randIndex] = q.problems[randIndex], q.problems[n-1]
		}
	}

	return nil
}
