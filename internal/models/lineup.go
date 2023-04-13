package models

type Lineup struct {
	ID        int    `json:"id"`
	Agent     string `json:"agent"`
	MapName   string `json:"map"`
	Objective bool   `json:"objective"`
	Ability   string `json:"ability"`
	Source    string `json:"source"`
}
