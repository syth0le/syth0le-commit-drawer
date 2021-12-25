package main

type ContributionTable struct {
	days  int8 // 7
	weeks int8 // 52 - 54
}

func getSizeOfContribution() ContributionTable {
	contrTable := ContributionTable{
		days:  7,
		weeks: 53,
	}

	return contrTable
}
