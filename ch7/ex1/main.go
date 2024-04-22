package main

import "fmt"

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func main() {
	realMadrid := Team{
		Name:    "Real Madrid",
		Players: []string{"Modric", "Carvajal"},
	}
	barcelona := Team{
		Name:    "Barcelona",
		Players: []string{"Gundhagen", "Pedri"},
	}
	spanishLeague := League{
		Teams: []Team{realMadrid, barcelona},
		Wins:  make(map[string]int),
	}
	for _, t := range spanishLeague.Teams {
		fmt.Println(t.Name)
		fmt.Println(t.Players)
	}
}
