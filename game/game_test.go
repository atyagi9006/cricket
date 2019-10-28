package game

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/atyagi9006/cricket/match"
)

func TestPlay(t *testing.T) {

	testCases := map[string]struct {
		Input    *match.Team
		Expected error
	}{
		"For a given nil team game should return Invalid input error ": {
			nil,
			ErrInvalidInput,
		},
		"For a given empty team game should return  error ": {
			&match.Team{},
			ErrInvalidTeamName,
		},
		"For a given team with name and without data game should return  error ": {
			&match.Team{
				TeamName: "Test",
			},
			ErrInvalidTeamData,
		},

		"for a given team target run greater than equal to 10  if not give Incorrect Teamdata Error ": {
			&match.Team{
				TeamName: "Test",
				Data: &match.Data{
					TargetRun: 9,
				},
			},
			ErrIncorrectTargetRuns,
		},
		"for a given team number of over should be greater than equal to 1  if not give Incorrect Teamdata Error ": {
			&match.Team{
				TeamName: "Test",
				Data: &match.Data{
					TargetRun: 10,
					Over:      0,
				},
			},
			ErrIncorrectOvers,
		},
		"for a given team number of players should be greater than equal to 2  if not give Incorrect Teamdata Error ": {
			&match.Team{
				TeamName: "Test",
				Data: &match.Data{
					TargetRun: 10,
					Over:      1,
					Players:   []string{"test1"},
				},
			},
			ErrIncorrectPlayers,
		},
		"for a given team number of players should be equal to the length of probability if not give Incorrect Teamdata Error ": {
			&match.Team{
				TeamName: "Test",
				Data: &match.Data{
					TargetRun: 10,
					Over:      1,
					Players:   []string{"test1", "test2", "test3", "test4"},
					PlayerScoreProbablity: map[string][]int{
						"test1": []int{5, 30, 25, 10, 15, 1, 9, 5},
						"test2": []int{5, 30, 25, 10, 15, 1, 9, 5},
						"test3": []int{5, 30, 25, 10, 15, 1, 9, 5},
					},
				},
			},
			ErrIncorrectTeamData,
		},
		"For a given team game should be played and result should come": {
			match.NewTeam(),
			nil,
		},
	}

	Convey("GamePlay Test", t, func() {
		for testcaseName, testCaseData := range testCases {
			Convey(testcaseName, func() {
				game := NewGame()
				result, err := game.Play(testCaseData.Input)
				So(err, ShouldEqual, testCaseData.Expected)
				if err == nil {
					So(result, ShouldNotBeNil)
					So(result.Message, ShouldContainSubstring, "team")
				}

			})
		}
	})

}
