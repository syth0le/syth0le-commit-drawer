package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func runGitStatus() {
	command := "status"
	output := execute(command)
	writeToLogFile(command)
	_ = output
	// writeToLogFile(output)
}

func runGitAdd() {
	command := []string{"add", "."}
	execute(command...)
	writeToLogFile(command...)
}

func runGitCommit(number string) {
	dateTime := getDateTimeCommit()
	message := getRandomCommitMessage(number)
	command := []string{"commit", "-m", message, "--date", dateTime}
	output := execute(command...)
	writeToLogFile(command...)
	writeToLogFile(output)
}

func runGitPush() {
	command := "push"
	execute(command)
	writeToLogFile(command)
}

func execute(command ...string) string {
	out, err := exec.Command("git", command...).Output()
	if err != nil {
		fmt.Printf("%s - git %s\n", err, strings.Join(command, " "))
	} else {
		fmt.Printf("Command Successfully Executed - git %s\n", strings.Join(command, " "))
	}

	output := string(out[:])
	return output
}

func runCommands(config ContributionTable) {
	table := strings.Split(config.table, "\n")[1:8]
	for i := 0; i < len(table[0]); i++ {
		for j := 0; j < config.days; j++ {
			symbol := table[j][i]
			number := getNumber()

			if string(symbol) == "#" {
				for k := 0; k < 5; k++ {
					numberToMessage := strconv.Itoa(number) + "-" + strconv.Itoa(k)
					writeChangesToFile(numberToMessage)
					runGitStatus()
					runGitAdd()
					runGitCommit(numberToMessage)
				}
			}
		}
	}
	runGitPush()
	cleanLogFile(false)
}
