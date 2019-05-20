package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Nachbarschaften(t *testing.T) {
	de := Staat{ID: "DE"}
	nl := Staat{ID: "NL"}
	for _, test := range []struct {
		Name            string
		Nachbarschaften Nachbarschaften
		Nachbarn        bool
	}{
		{
			Name: "NullValue",
		},
		{
			Name:            "Nachbarn",
			Nachbarschaften: Nachbarschaften{"DE": map[string]bool{"NL": true}},
			Nachbarn:        true,
		},
		{
			Name:            "Nachbarn other way around",
			Nachbarschaften: Nachbarschaften{"NL": map[string]bool{"DE": true}},
			Nachbarn:        true,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Nachbarn, test.Nachbarschaften.SindNachbarn(de, nl))
		})
	}
}
