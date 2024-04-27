package main

import (
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
)

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for i, r := range r.Ranking() {
		io.WriteString(w, strconv.Itoa(i+1)+" "+r+"\n")
	}
}

type MatchResult struct {
	t1 string
	s1 int
	t2 string
	s2 int
}

type Team struct {
	Name    string
	Players []string
}

type Pair struct {
	Key string
	Val int
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) Ranking() []string {
	winners := []string{}
	pairs := []Pair{}
	for k, v := range l.Wins {
		pairs = append(pairs, Pair{Key: k, Val: v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Val > pairs[j].Val })
	for _, p := range pairs {
		winners = append(winners, p.Key)
	}
	return winners
}

func (l League) TeamExists(team string) bool {
	found := false
	for _, t := range l.Teams {
		if t.Name == team {
			found = true
			break
		}
	}
	return found
}

func (l *League) AddMatchResult(mr MatchResult) error {
	if !l.TeamExists(mr.t1) {
		return errors.New(mr.t1 + " does not exist")
	}
	if !l.TeamExists(mr.t2) {
		return errors.New(mr.t2 + " does not exist")
	}
	if mr.s1 > mr.s2 {
		l.Wins[mr.t1]++
	}
	if mr.s2 > mr.s1 {
		l.Wins[mr.t2]++
	}
	return nil
}

func main() {

	teamNames := []string{"Liverpool", "Chelsea", "Newcastle", "Brentford"}
	playerNames := []string{"Haaland", "Iwobi", "Aguero", "Pedro"}
	matchResults := []MatchResult{
		MatchResult{t1: "Liverpool", s1: 3, t2: "Chelsea", s2: 2},
		MatchResult{t1: "Liverpool", s1: 5, t2: "Newcastle", s2: 1},
		MatchResult{t1: "Brentford", s1: 2, t2: "Liverpool", s2: 0},
		MatchResult{t1: "Newcastle", s1: 1, t2: "Chelsea", s2: 1},
	}
	league := League{
		Teams: []Team{},
		Wins:  make(map[string]int),
	}

	for _, t := range teamNames {
		league.Teams = append(league.Teams, Team{Name: t, Players: playerNames})
		league.Wins[t] = 0
	}
	for _, mr := range matchResults {
		league.AddMatchResult(mr)
	}

	RankPrinter(league, os.Stdout)
}
