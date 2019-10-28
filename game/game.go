package game

import (
	"errors"
	"fmt"
	"log"

	"github.com/atyagi9006/cricket/match"
	"github.com/atyagi9006/cricket/random"
)

var (
	// ErrInvalidInput ...
	ErrInvalidInput = errors.New("input for team can't be nil")
	// ErrInvalidTeamName ...
	ErrInvalidTeamName = errors.New("Team name can't be empty")
	// ErrInvalidTeamData ...
	ErrInvalidTeamData = errors.New("Team Data can't be nil")
	// ErrIncorrectPlayers ...
	ErrIncorrectPlayers = errors.New("Team Data number of players should be greater than equal to 2")
	// ErrIncorrectTeamData ...
	ErrIncorrectTeamData = errors.New("Team Data number of players should be equal to len of player score probability")
	// ErrIncorrectTargetRuns ...
	ErrIncorrectTargetRuns = errors.New("Team Data target run should be greater than equal to 10")
	// ErrIncorrectOvers ...
	ErrIncorrectOvers = errors.New("Team Data overs to play should be greater than equal to 1")
)

// Game ...
type Game struct{}

// PlayerStatus ...
type PlayerStatus struct {
	Score int
	Balls int
	Out   bool
}

// Result ... contains the final result
type Result struct {
	Message     string
	FinalScores map[string]PlayerStatus
}

// NewGame ...
func NewGame() Game {
	return Game{}
}

// Play will start game play
func (g Game) Play(team *match.Team) (*Result, error) {
	if team == nil {
		return nil, ErrInvalidInput
	}

	if team.TeamName == "" {
		return nil, ErrInvalidTeamName
	}

	if team.Data == nil {
		return nil, ErrInvalidTeamData
	}

	if team.Data.TargetRun < 10 {
		return nil, ErrIncorrectTargetRuns
	}

	if team.Data.Over < 1 {
		return nil, ErrIncorrectOvers
	}

	if len(team.Data.Players) < 2 {
		return nil, ErrIncorrectPlayers
	}

	if len(team.Data.Players) != len(team.Data.PlayerScoreProbablity) {
		return nil, ErrIncorrectTeamData
	}

	runs := team.Data.TargetRun
	overs := team.Data.Over
	wickets := team.Data.Wicket
	players := team.Data.Players
	remaining := players[2:]
	// Scores for players, with an out flag
	//If the player hasn't been on the field yet, their scores will not be present here
	scores := map[string]PlayerStatus{
		players[0]: {Score: 0, Balls: 0, Out: false},
		players[1]: {Score: 0, Balls: 0, Out: false},
	}

	probs := team.Data.PlayerScoreProbablity
	//  Players on  strike (1st and 2nd on the field)
	playing := []string{players[wickets], players[wickets+1]}

	log.Println("Commentry ....")
	for over := 0; over < overs; over++ {
		printOverStats(overs-over, runs)
		for ball := 0; ball < match.Balls; ball++ {
			randomRun, err := random.Generator(probs[playing[0]])
			if err != nil {
				return nil, err
			}
			if score, ok := scores[playing[0]]; ok {
				score.Balls++
				scores[playing[0]] = score
			}
			if randomRun != 7 {
				// Reducing number of runs remaining
				runs = runs - randomRun
				// Increasing the score of the player
				if score, ok := scores[playing[0]]; ok {
					score.Score += randomRun
					scores[playing[0]] = score
				}
				// Print the score for that ball
				printBallStats(over, ball, randomRun, playing[0], false)
				// If the no. of runs is odd then change strike
				if randomRun%2 != 0 {
					changestrike(playing)
				}

				//More than given target runs made, team wins
				if runs <= 0 {
					message := fmt.Sprintln(team.TeamName, " team won by ", (4 - wickets), " wickets and ", ((overs-1-over)*6)+
						(5-ball), " balls remaining")
					result := Result{
						Message:     message,
						FinalScores: scores,
					}
					return &result, nil
				}

			} else {
				//If randno is 7 and the player is out
				wickets++
				// Set the player status to Out
				if score, ok := scores[playing[0]]; ok {
					score.Out = true
					scores[playing[0]] = score
				}
				printBallStats(over, ball, 0, playing[0], true)

				// If 3 players are out, team lost
				if wickets == 3 {
					message := fmt.Sprintln(team.TeamName, "team lost by ", runs, " runs")
					result := Result{
						Message:     message,
						FinalScores: scores,
					}
					return &result, nil
				}
				// Put the next player on strike
				playing[0] = remaining[0]
				scores[remaining[0]] = PlayerStatus{
					Score: 0,
					Balls: 0,
					Out:   false,
				}
				//Remove onstrike player from remaining players list
				removeRemainingPlayer(remaining, 0)

			}

		}
		changestrike(playing)
	}

	//     If all overs get finished and target run also equal to 0 then the game is tie down
	if runs == 0 {
		log.Println("Match tied!")
	}
	// If runs scored are less than 40 and balls are over
	message := fmt.Sprintln(team.TeamName, " team lost by ", runs, " runs")
	result := Result{
		Message:     message,
		FinalScores: scores,
	}
	return &result, nil
}

func printOverStats(overs, runs int) {
	log.Println(overs, " overs left. ", runs, " runs to win")
}

func printBallStats(over, ball, score int, player string, out bool) {
	if out {
		log.Println(over, ".", ball, " ", player, " Out!")
	} else {
		log.Println(over, ".", (ball + 1), " ", player, " scores ", score, " run")
	}

}

func changestrike(playing []string) {
	playing[0], playing[1] = playing[1], playing[0]
}

//PrintScores TO PRINT Final result
func (g Game) PrintScores(result *Result) {
	log.Println(" Result: ", result.Message)
	log.Println("----------------- SCOREBOARD -----------------")
	for player, score := range result.FinalScores {
		//If player was not out then we need to add a "*" after the score
		if score.Out {
			log.Println(player, " - ", score.Score, " (", score.Balls, " balls)")
		} else {
			log.Println(player, " - ", score.Score, "* (", score.Balls, " balls)")
		}
	}

}

func removeRemainingPlayer(remaining []string, i int) {
	copy(remaining[i:], remaining[i+1:]) // Shift remaining[i+1:] left one index.
	remaining[len(remaining)-1] = ""     // Erase last element (write zero value).
	remaining = remaining[:len(remaining)-1]
}
