package main

import "fmt"

func getDateTimeCommit(date string) string {
	return date
}

func getRandomCommitMessage(number int) string {
	// randomize function
	message := "title message"
	return message
}

func runGitAdd() {
	command := "git add ."
	fmt.Println(command)
}

func runGitCommit() {
	dateTime := getDateTimeCommit("s")
	message := getRandomCommitMessage(1)
	command := "git commit -m " + message + dateTime
	fmt.Println(command)
}

func runGitPush() {
	command := "git push"
	fmt.Println(command)
}

func runCommands() {
	runGitAdd()
	runGitCommit()
	runGitPush()
}
