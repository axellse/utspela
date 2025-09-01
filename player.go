package main

import (
	"bytes"
	"os/exec"
)

type Player interface {
	//Play plays a piece of raw audio, returning when the playback completes or when an error is encountered
	Play(a Audio) error
}

type StdinPlayer struct{
	Command string
	Args []string
}
func (s StdinPlayer) Play(a Audio) error {
	cmd := exec.Command(s.Command, s.Args...)
	cmd.Stdin = bytes.NewReader(a)
	return cmd.Run()
}

var Players = map[string]Player{
	"ffplay" : StdinPlayer{
		Command: "ffplay",
		Args: []string{"-", "-nodisp", "-autoexit"},
	},
	"sox-mp3" : StdinPlayer{
		Command: "play",
		Args: []string{"-t", "mp3", "-"},
	},
}
var CurrentPlayer Player