package codewars

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type User struct {
	Username            string         `json:"username"`
	Name                string         `json:"name"`
	Honor               int            `json:"honor"`
	Clan                string         `json:"clan"`
	LeaderboardPosition int            `json:"leaderboardPosition"`
	Ranks               Ranks          `json:"ranks"`
	CodeChallenges      CodeChallenges `json:"codeChallenges"`
}
type Ranks struct {
	Overall   Overall   `json:"overall"`
	Languages Languages `json:"languages"`
}
type Overall struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Languages map[string]Language

type Language struct {
	Score int `json:"score"`
}

type CodeChallenges struct {
	TotalCompleted int `json:"totalCompleted"`
}

var LevelColors = map[string]string{
	"1 kyu": "#866CC7",
	"2 kyu": "#866CC7",
	"3 kyu": "#3C7EBB",
	"4 kyu": "#3C7EBB",
	"5 kyu": "#ECB613",
	"6 kyu": "#ECB613",
	"7 kyu": "#E6E6E6",
	"8 kyu": "#E6E6E6",
	"1 dan": "#999999",
	"2 dan": "#999999",
}

func (u *User) GetUserData(username string) error {
	resp, err := http.Get("https://www.codewars.com/api/v1/users/" + username)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("user not found")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &u)
	if err != nil {
		return err
	}
	return nil
}
