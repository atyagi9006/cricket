package match

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Balls represents balls in a over
const Balls = 6

// Team represents Cricket team...
type Team struct {
	TeamName string `yaml:"team_name"`
	Data     *Data  `yaml:"data"`
}

// Data struct represents Cricket Team Data
type Data struct {
	TargetRun             int `yaml:"target_run"`
	Over                  int `yaml:"over"`
	Wicket                int
	Players               []string         `yaml:"players"`
	Score                 []string         `yaml:"score"`
	PlayerScoreProbablity map[string][]int `yaml:"player_score_probablity"`
}

// NewTeam func returns a new team pointer every time
func NewTeam() *Team {
	log.Println("Loading Team")
	t := &Team{}
	return t.LoadTeam()
}

// LoadTeam func load team from a team ymal
func (t *Team) LoadTeam() *Team {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/config/bengaluru.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return t
}
