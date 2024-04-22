package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(t1, t2 string, s1, s2 int) {
	if s1 > s2 {
		w, ok := l.Wins[t1]
		if ok {
			l.Wins[t1] = w + 1
		} else {
			l.Wins[t1] = 1
		}
	} else if s2 > s1 {
		w, ok := l.Wins[t2]
		if ok {
			l.Wins[t2] = w + 1
		} else {
			l.Wins[t2] = 1
		}
	}
}

func (l League) Ranking() []string {
	result := make([]string, 0, 0)
	scoreMap := make(map[int][]string)
	scores := make([]int, 0, 0)
	for k, v := range l.Wins {
		teamNames, ok := scoreMap[v]
		if ok {
			containsTeam := false
			for _, t := range teamNames {
				if t == k {
					containsTeam = true
					break
				}
			}
			if !containsTeam {
				teamNames = append(teamNames, k)
			}
		} else {
			scoreMap[v] = []string{k}
		}
		scores = append(scores, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	for _, s := range scores {
		teams := scoreMap[s]
		for _, t := range teams {
			r := fmt.Sprintf("Team[%s]Wins[%d]", t, s)
			result = append(result, r)
		}
	}
	return result
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
	laLiga := League{
		Teams: []Team{realMadrid, barcelona},
		Wins:  make(map[string]int),
	}
	laLiga.MatchResult("Real Zarazoga", "Barcelona", 3, 2)
	laLiga.MatchResult("Real Madrid", "Barcelona", 3, 1)
	laLiga.MatchResult("Benfica", "Barcelona", 1, 3)
	laLiga.MatchResult("Real Madrid", "Barcelona", 2, 5)
	laLiga.MatchResult("Real Madrid", "Getafe", 1, 2)
	laLiga.MatchResult("Real Madrid", "Barcelona", 3, 2)
	laLiga.MatchResult("Getafe", "Barcelona", 3, 1)
	laLiga.MatchResult("Real Madrid", "Barcelona", 1, 3)
	laLiga.MatchResult("Real Madrid", "Atletico Madrid", 2, 5)
	laLiga.MatchResult("Real Betis", "Barcelona", 1, 2)
	laLiga.MatchResult("Real Madrid", "Barcelona", 3, 2)
	laLiga.MatchResult("Getafe", "Barcelona", 3, 1)
	laLiga.MatchResult("Real Madrid", "Mallorca", 1, 3)
	laLiga.MatchResult("Real Sociodad", "Barcelona", 2, 5)
	laLiga.MatchResult("Real Madrid", "Barcelona", 1, 2)
	rankings := laLiga.Ranking()
	for _, r := range rankings {
		fmt.Println(r)
	}
}
