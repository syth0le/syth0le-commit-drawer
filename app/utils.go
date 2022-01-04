package main

import (
	"os"
	"time"
)

func getNumber() int {
	globalAccumulator++
	return globalAccumulator
}

func getDateTimeCommit() string {
	config := getSizeOfContribution()
	difference := config.wholeYear - globalAccumulator + 1
	dateTime := time.Now()
	then := dateTime.Add(time.Duration(-24*difference) * time.Hour)
	return then.Format("2006-01-02T15:04:05-0700")
}

func getRandomCommitMessage(number string) string {
	message := "'" + "title message #" + number + "'"
	return message
}

func writeChangesToFile(number string) bool {
	isWrited := false
	file, err := os.OpenFile("changeable_file.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	valueToWrite := "Hello" + number + "\n"

	_, err = file.WriteString(valueToWrite)
	if err != nil {
		panic(err)
	}
	file.Close()
	isWrited = true
	return isWrited
}
