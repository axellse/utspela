package main

import (
	"math/rand"
	"os"
	"time"
)

func PlayComponent(c BroadcastComponent) {
	switch c.Type {
	case "file":
		ba, ferr := os.ReadFile(c.MediaSource)
		Info("Starting next broadcast segment at " + time.Now().Format(time.TimeOnly))
		if ferr != nil {
			Error("Could not read media file, skipping segment")
			return
		}
		err := CurrentPlayer.Play(ba)
		if err != nil {
			Error("Encountered playback error, skipping segment")
			return
		}
	case "rand":
		ci := rand.Intn(len(c.SubComponents) - 1)
		comp := c.SubComponents[ci]
		PlayComponent(comp)
	}

	Done("Broadcast segment complete at " + time.Now().Format(time.TimeOnly))
}

var BroadcastActive bool

func BeginBroadcast(b Broadcast) {
	if BroadcastActive {
		Info("A broadcast should begin now, but one is already airing, cancelling this one.")
		return
	}
	Info("Starting broadcast at " + time.Now().Format(time.TimeOnly))
	BroadcastActive = true
	for _, c := range b.Components {
		PlayComponent(c)
	}
	BroadcastActive = false
	Done("Broadcast complete on " + time.Now().Format(time.TimeOnly))
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