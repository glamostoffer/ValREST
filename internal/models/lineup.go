package models

type Lineup struct {
	ID        int    `json:"id"`
	Agent     string `json:"agent"`
	MapName   string `json:"map"`
	Objective string `json:"objective"`
	Ability   string `json:"ability"`
	Source    string `json:"source"`
	Author    int    `json:"author"` // id пользователя, создавшего lineup
}
