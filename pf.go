package main

import (
	"encoding/json"
)

var ProgrammingFile []Broadcast

type BroadcastComponent struct {
	Type          string               `json:"type"`
	MediaSource   string               `json:"src"`
	SubComponents []BroadcastComponent `json:"components"`
}

type Broadcast struct {
	At         string               `json:"at"`
	Components []BroadcastComponent `json:"components"`
	asttime AstTime
}

func ParseProgrammingFile(pf []byte) {
	err := json.Unmarshal(pf, &ProgrammingFile)
	if err != nil {
		FError("json unmarshal programming file error: " + err.Error())
	}

	newFile := []Broadcast{}
	for _, v := range ProgrammingFile {
		v.asttime, err = ParseString(v.At)
		if err != nil {
			FError("invalid at value in programming file: " + err.Error())
		}
		newFile = append(newFile, v)
	}

	ProgrammingFile = newFile
	Done("Read programming")
}