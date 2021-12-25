package main

var globalAccumulator int = 0

type ContributionTable struct {
	days  int // 7
	weeks int // 52 - 54
}

func getSizeOfContribution() ContributionTable {
	contrTable := ContributionTable{
		days:  7,
		weeks: 53,
	}

	return contrTable
}
