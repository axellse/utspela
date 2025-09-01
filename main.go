package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(Green + "Utspela" + Reset + " is audio playout software, the name comes from the word 'playout' in swedish.")
	fmt.Println("Running version " + Green + "1.0.0" + Reset + ", written by axell (axell.me)")
	fmt.Println()
	Info("Utspela has been launched!")

	sp := flag.String("pf", "", "a path to a programming file")
	player := flag.String("p", "ffplay", "type of player to use, options are: 'ffplay' (default), 'sox-mp3'")

	flag.Parse()
	if *sp == "" {
		FError("please provide a programming file!")
	}
	if plr, ok := Players[*player]; ok {
		CurrentPlayer = plr
	} else {
		FError("invalid player provided!")
	}
	
	pf, err := os.ReadFile(*sp)
	if err != nil {
		FError("read programming file error: " + err.Error())
	}
	ParseProgrammingFile(pf)

	ListenForBroadcasts()

}