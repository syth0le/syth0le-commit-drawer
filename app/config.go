package main

import (
	"fmt"
	"time"
)

// NOT CHANGEABLE
var globalAccumulator int = 1
var daysOfWeekDifference = getDaysOfWeekDifference()

// CHANGEABLE
var isWriteToLog bool = false
var isCleanLogFile bool = false
var commitNumbersPerDay = 5

type ContributionTable struct {
	days      int // 7
	weeks     int // 52 - 54
	wholeYear int // 365 - 366
	table     string
}

var table string = `
----###-#-###-###--###-#-###--#####-###-###-####---
----#-#-#-#-#-#----#-#-#-#-#----#---#---#---#--#---
----#-#-#-#-#-#----#-#-#-#-#----#---#---#---#--#---
----#-#-###-#-###--#-#-###-#----#---###-###-####---
----#-#-#-#-#-#-#--#-#-#-#-#----#---#---#-#---##---
----#-#-#-#-#-#-#--#-#-#-#-#----#---#---#-#--#-#---
---##-#-#-###-###-##-#-#-###----#---###-###-#--#---
`

// или сделать вычисляемой исходя из дня ()
// проверить на отрисовке сначала (не в гите а на пикче)
func getSizeOfContribution() ContributionTable {
	contrTable := ContributionTable{
		days:      7,
		weeks:     53,
		wholeYear: 366,
		table:     table,
	}

	return contrTable
}

func getDaysOfWeekDifference() int {
	daysOfWeek := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}
	_ = daysOfWeek
	day := time.Now().Weekday().String()
	fmt.Printf("%s", day)
	return daysOfWeek[day]

}
