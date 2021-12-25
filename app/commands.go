package main

import "fmt"

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
	writeChangesToFile(2)
	runGitAdd()
	runGitCommit()
	runGitPush()
}
