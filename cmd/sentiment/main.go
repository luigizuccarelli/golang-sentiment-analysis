package main

import (
	"fmt"

	"github.com/luigizuccarelli/golang-sentiment-analysis/pkg/sentiment"
)

func main() {
	fmt.Println("Executing training on english language")
	//model, err := sentiment.Train(os.Args[1])
	model, err := sentiment.Restore()
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
