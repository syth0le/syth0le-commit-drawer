package main

var globalAccumulator int = 1
var isWriteToLog bool = false
var isCleanLogFile bool = false

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
