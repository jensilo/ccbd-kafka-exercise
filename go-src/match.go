package src

type Score struct {
	Team1     string `json:"team1"`
	Team2     string `json:"team2"`
	Score1    int    `json:"score1"`
	Score2    int    `json:"score2"`
	TimeStamp string `json:"timestamp"`
}
