package main

import (
	"os"
	"strconv"
	"time"
)

func getNumber() int {
	globalAccumulator++
	return globalAccumulator
}

func getDateTimeCommit() string {
	config := getSizeOfContribution()
	difference := config.wholeYear - globalAccumulator
	dateTime := time.Now()
	then := dateTime.Add(time.Duration(-24*difference) * time.Hour)
	return then.Format("2006-01-02T15:04:05-0700")
}

func getRandomCommitMessage(number int) string {
	message := "title message #" + strconv.Itoa(number)
	return message
}

func writeChangesToFile(number int) bool {
	isWrited := false
	file, err := os.OpenFile("changeable_file.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	valueToWrite := "Hello" + strconv.Itoa(number) + "\n"

	_, err = file.WriteString(valueToWrite)
	if err != nil {
		isWrited = true
		panic(err)
	}
	file.Close()

	return isWrited
}
