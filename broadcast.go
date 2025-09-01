package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PlayComponent(c BroadcastComponent) {
	if c.Type == "file" {
		ba, ferr := os.ReadFile(c.MediaSource)
		if ferr != nil {
			fmt.Println("broadcast play failiure, skipping...")
			return
		}
		err := CurrentPlayer.Play(ba)
		if err != nil {
			fmt.Println("playback error, skipping...")
		}
	} else if c.Type == "rand" {
		ci := rand.Intn(len(c.SubComponents) - 1)
		comp := c.SubComponents[ci]
		PlayComponent(comp)
	}
}

var BroadcastActive bool

func BeginBroadcast(b Broadcast) {
	if BroadcastActive {
		Info("A broadcast should begin now, but one is already airing, cancelling this one.")
		return
	}
	Info("Starting broadcast on " + time.Now().Format(time.TimeOnly))
	BroadcastActive = true
	for _, c := range b.Components {
		PlayComponent(c)
	}
	BroadcastActive = false
	Info("Broadcast complete on " + time.Now().Format(time.TimeOnly))
}

func ListenForBroadcasts() {
	startTime := time.Now()
	untilNextWhole := time.Second - time.Duration(startTime.Nanosecond())
	time.Sleep(untilNextWhole)
	
	ticker := time.NewTicker(time.Second)
	Info("Ready to begin broadcasting")
	for {
		<- ticker.C
		for _, v := range ProgrammingFile {
			if v.asttime.Ongoing() {
				go BeginBroadcast(v)
			}
		}
	}
}