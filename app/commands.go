package main

import (
	"fmt"
	"os/exec"
)

func runGitStatus() {
	command := "git status"
	output := execute(command)
	writeToLogFile(command)
	writeToLogFile(output)
}

func runGitAdd() {
	command := "git add ."
	writeToLogFile(command)
}

func runGitCommit(number int) {
	dateTime := getDateTimeCommit()
	message := getRandomCommitMessage(number)
	command := "git commit -m " + message + " --date=" + dateTime
	writeToLogFile(command)
}

func runGitPush() {
	command := "git push"
	writeToLogFile(command)
}

func execute(command string) string {
	out, err := exec.Command("cmd", "/C", command).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	return output
}

func runCommands() {
	for globalAccumulator < 10 {
		number := getNumber()
		writeChangesToFile(number)
		runGitStatus()
		runGitAdd()
		runGitCommit(number)
		runGitPush()
	}
	cleanLogFile(false)
}
