package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"ep1-quizgame/quiz"
)

var (
	fileName  = flag.String("file", "problems.csv", "-file=path/to/file")
	timeLimit = flag.Int("time", 30, "-time=time limit (seconds)")
	shuffle   = flag.Bool("shuffle", false, "-shuffle  Shuffle questions")
)

func main() {

	flag.Parse()

	q := quiz.MakeQuiz(*fileName, *shuffle)
	q.ReadQuiz()

	fmt.Printf("\n\nWelcome to the Quiz, You have %d seconds to complete. Results will be provided immediatly upon finish\n\n", *timeLimit)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(*timeLimit))
	defer cancel()

	go q.BeginQuiz()

	select {
	case <-ctx.Done():
		fmt.Println("Context done")
	case <-time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Printf("\n\n!! Time Up !!\n")
		fmt.Printf("Final Score %d/%d", q.GetUserScore(), q.GetTotalQuestios())
	}
}
