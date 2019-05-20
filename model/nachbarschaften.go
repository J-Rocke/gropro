package model

type Nachbarschaften map[string]map[string]bool

func (n Nachbarschaften) SindNachbarn(a, b Staat) bool {
	return n[a.ID][b.ID] || n[b.ID][a.ID]
}
