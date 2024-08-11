package main

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

/*
1. You have been asked to manage a basketball league and are going to write a program
to help you.
Define two types. The first one, called Team, has a field for the name of the
team and a field for the player names. The second type is called League and has
a field called Teams for the teams in the league and a field called Wins that
maps a team’s name to its number of wins.

2. Add two methods to League. The first method is called MatchResult.
It takes four parameters: the name of the first team, its score in the game,
the name of the second team, and its score in the game.
This method should update the Wins field in League.
Add a second method to League called Ranking that returns a slice of the
team names in order of wins.
Build your data structures and call these methods from the main function
in your program using some sample data.

3. Define an interface called Ranker that has a single method called
Ranking that returns a slice of strings. Write a function called
RankPrinter with two parameters, the first of type Ranker and the
second of type io.Writer. Use the io.WriteString function to write
the values returned by Ranker to the io.Writer, with a newline separating
each result. Call this function from main.
*/

type Team struct {
	Name        string   `json:"teamName"`
	PlayerNames []string `json:"players"`
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) AddMatchResult(t1 string, t1Score int, t2 string, t2Score int) {
	winner, err := l.GetWinner(t1, t1Score, t2, t2Score)
	if err == nil {
		wins, ok := l.Wins[winner]
		if !ok {
			l.Wins[winner] = 1
		} else {
			l.Wins[winner] = wins + 1
		}
	}
}

func (l League) GetWinner(t1 string, t1Score int, t2 string, t2Score int) (string, error) {
	if t1Score > t2Score {
		return t1, nil
	} else if t2Score > t1Score {
		return t2, nil
	} else {
		return "", errors.New("no winner")
	}
}

func (l League) GetRanking() []string {
	winsSlice := make([][]string, 0, 5)
	for k, v := range l.Wins {
		winsSlice = append(winsSlice, []string{strconv.Itoa(v), k})
	}
	slices.SortFunc(winsSlice, func(a, b []string) int {
		aScore, err := strconv.Atoi(a[0])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		bScore, err := strconv.Atoi(b[0])
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return cmp.Compare(bScore, aScore)
	})
	resultSlice := make([]string, 0, 5)
	for _, w := range winsSlice {
		resultSlice = append(resultSlice, w[1]+": "+w[0])
	}
	return resultSlice
}

func GenerateLeague() (League, error) {
	league := League{
		Teams: make([]Team, 0, 5),
		Wins:  make(map[string]int),
	}
	jsonFile, err := os.Open("data.json")
	if err != nil {
		return league, err
	}
	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return league, err
	}
	teams := make([]Team, 0, 5)
	err = json.Unmarshal(jsonBytes, &teams)
	if err != nil {
		return league, err
	}
	for _, t := range teams {
		league.Teams = append(league.Teams, t)
		league.Wins[t.Name] = 0
	}
	return league, nil
}

func main() {
	league, err := GenerateLeague()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	league.AddMatchResult("Real Madrid", 3, "Barcelona", 2)
	league.AddMatchResult("Bayern Munich", 1, "Barcelona", 2)
	league.AddMatchResult("Paris Saint-Germain", 0, "Barcelona", 2)
	league.AddMatchResult("Real Madrid", 1, "Manchester City", 5)
	for _, rank := range league.GetRanking() {
		fmt.Println(rank)
	}
}
