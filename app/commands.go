package main

import (
	"fmt"
	"os/exec"
	"strings"
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
	config := getSizeOfContribution()
	table := strings.Split(config.table, "\n")[1:8]
	for i := 0; i < len(table); i++ {
		for j := 0; j < config.days; j++ {

			// fmt.Println(j, i)
			// fmt.Println(table[j])
			symbol := table[j][i]
			// fmt.Println(string(symbol))
			number := getNumber()
			if string(symbol) == "#" {
				writeChangesToFile(number)
				runGitStatus()
				runGitAdd()
				runGitCommit(number)
				runGitPush()
			}
		}
	}
	cleanLogFile(false)
}
