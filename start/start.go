package start

import (
	"log"

	"github.com/atyagi9006/cricket/game"
	"github.com/atyagi9006/cricket/match"
)

// Run ...
func Run() {
	log.Println("Starting Game ...")
	strikeTeam := match.NewTeam()
	g := game.NewGame()
	result, err := g.Play(strikeTeam)
	if err != nil {
		log.Println("Error in Starting cricket game", err)
	}
	g.PrintScores(result)

}
