/*
  Given an array of pairs representing the teams that have competed against each
  other and an array containing the results of each competition, write a
  function that returns the winner of the tournament. The input arrays are named
  <span>competitions</span> and <span>results</span>, respectively. The
  <span>competitions</span> array has elements in the form of
  <span>[homeTeam, awayTeam]</span>, where each team is a string of at most 30
  characters representing the name of the team. The <span>results</span> array
  contains information about the winner of each corresponding competition in the
  <span>competitions</span> array. Specifically, <span>results[i]</span> denotes
  the winner of <span>competitions[i]</span>, where a <span>1</span> in the
  <span>results</span> array means that the home team in the corresponding
  competition won and a <span>0</span> means that the away team won.

  It's guaranteed that exactly one team will win the tournament and that each
  team will compete against all other teams exactly once. It's also guaranteed
  that the tournament will always have at least two teams.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	results := []int{0, 1, 1}

	fmt.Println(TournamentWinner(competitions, results))
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {

	resultados := make(map[string]int)
	currentWinner := ""
	currentMaxScore := 0

	for i, arr := range competitions {
		winnerTeam := int(math.Abs(float64(results[i] - 1)))
		resultados[arr[winnerTeam]] += 1
		if resultados[arr[winnerTeam]] > currentMaxScore {
			currentWinner = arr[winnerTeam]
			currentMaxScore = resultados[arr[winnerTeam]]
		}
	}

	return currentWinner
}
