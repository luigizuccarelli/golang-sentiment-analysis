package sentiment

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

// TrainModel takes in a path to the expected
// datasets, from the file training.1600000.processed.noemoticon.csv
// and a map of models to add the model to. It'll return any errors if there were any.
func TrainModel(file string, modelMap Models) error {

	// open file for training
	// read line by line
	f, err := os.Open(file)

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := csv.NewReader(f)

	var class uint8
	var count int = 0

	stream := make(chan base.TextDatapoint, 1000)
	errors := make(chan error, 100)
	model := text.NewNaiveBayes(stream, 2, base.OnlyWords)

	go model.OnlineLearn(errors)

	for {
		data, err := scanner.Read()
		if err == io.EOF {
			break
		}
		// fmt.Println("LMZ DEBUG " + data[0] + " : " + data[5])
		if data[0] == "4" {
			class = 1
		} else {
			class = 0
		}
		stream <- base.TextDatapoint{
			X: data[5],
			Y: class,
		}
		if count%500 == 0 {
			print(".")
		}
		count++
	}

	now := time.Now()
	fmt.Printf("\nStart munging data from %s to model %v at %v\n", file, model, now)

	close(stream)

	for {
		err, more := <-errors
		if more {
			fmt.Printf("Error passed: %v\n", err)
		} else {
			// training is done!
			break
		}
	}

	modelMap[English] = model

	fmt.Printf("\nFinished munging from %s to data model\n\tdelta: %v\n\taverage time per line: %v\n", file, time.Now().Sub(now), time.Now().Sub(now)/time.Duration(int64(count)))
	return nil
}
