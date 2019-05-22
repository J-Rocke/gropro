package model

type AusgangsdatenContainer struct {
	Titel           string
	Staaten         []Staat
	Nachbarschaften Nachbarschaften
}

type LoesungsContainer struct {
	Titel     string
	Staaten   []Staat
	Iteration int
}
