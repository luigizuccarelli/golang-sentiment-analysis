package main

import (
	"fmt"
	"os"

	"github.com/luigizuccarelli/golang-sentiment-analysis/pkg/sentiment"
)

func main() {
	fmt.Println("Executing " + os.Args[1] + " on english language")

	var model sentiment.Models
	var err error

	if os.Args[1] == "training" {
		model, err = sentiment.Train(os.Args[2])
	} else {
		model, err = sentiment.Restore()
	}

	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	analysis := model.SentimentAnalysis("swine pig thats what you are", sentiment.English)
	fmt.Println("Analysis : ", analysis.Score)
	analysis = model.SentimentAnalysis("Its a lousy day", sentiment.English)
	fmt.Println("Analysis : ", analysis.Score)
	analysis = model.SentimentAnalysis("Its a beautiful day", sentiment.English)
	fmt.Println("Analysis : ", analysis.Score)

}
