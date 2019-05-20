package model

type AusgangsdatenContainer struct {
	Title           string
	Staaten         []Staat
	Nachbarschaften Nachbarschaften
}

type LoesungsContainer struct {
	Title string
	Staaten []Staat
}