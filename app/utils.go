package main

import (
	"os"
	"strconv"
)

func getDateTimeCommit(date string) string {
	return date
}

func getRandomCommitMessage(number int) string {
	// randomize function
	message := "title message"
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

	return isWrited
}
