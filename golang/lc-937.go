package main

import (
	"sort"
	"strings"
)

type Log struct {
	Index      int
	Identifier string
	Content    string
}

func reorderLogFiles(logs []string) []string {
	newLetterLogs := make([]Log, 0, len(logs))
	newDigitalLogs := make([]Log, 0, len(logs))
	for i := range logs {
		log := strings.Split(logs[i], " ")
		identifier := log[0]
		content := strings.Join(log[1:], " ")
		if content[0] >= '0' && content[0] <= '9' {
			newDigitalLogs = append(newDigitalLogs, Log{Index: i})
		} else {
			newLetterLogs = append(newLetterLogs, Log{Index: i, Identifier: identifier, Content: content})
		}
	}
	sort.Slice(newLetterLogs, func(i, j int) bool {
		if newLetterLogs[i].Content == newLetterLogs[j].Content {
			return newLetterLogs[i].Identifier < newLetterLogs[j].Identifier
		}
		return newLetterLogs[i].Content < newLetterLogs[j].Content
	})
	res := make([]string, 0, len(logs))
	for i := 0; i < len(newLetterLogs); i++ {
		res = append(res, logs[newLetterLogs[i].Index])
	}
	for i := range newDigitalLogs {
		res = append(res, logs[newDigitalLogs[i].Index])
	}
	return res
}
