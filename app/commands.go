package main

import "fmt"

func runGitAdd() {
	command := "git add ."
	fmt.Println(command)
}

func runGitCommit(number int) {
	dateTime := getDateTimeCommit("s")
	message := getRandomCommitMessage(number)
	command := "git commit -m " + message + " --date=" + dateTime
	fmt.Println(command)
}

func runGitPush() {
	command := "git push"
	fmt.Println(command)
}

func runCommands() {
	for globalAccumulator < 5 {
		number := getNumber()
		writeChangesToFile(number)
		runGitAdd()
		runGitCommit(number)
		runGitPush()
	}
}
