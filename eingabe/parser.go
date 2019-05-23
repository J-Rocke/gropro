package eingabe

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/J-Rocke/gropro/model"
)

// ParseFromFile liest Staaten und Nachbarschaften aus einer Datei
func ParseFromFile(filename string) (*model.AusgangsdatenContainer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return Parse(file)
}

// Parse liest Staaten und Nachbarschaften und schreibt sie in einen AusgangsdatenContainer
func Parse(r io.Reader) (*model.AusgangsdatenContainer, error) {
	container := model.AusgangsdatenContainer{}
	scanner := bufio.NewScanner(r)

	// Erste Line ist der Titel
	scanner.Scan()
	container.Titel = scanner.Text()

	var line string
	// Read all Staaten
	for scanner.Scan() {
		line = scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // Ignore comments
		}
		if strings.Contains(line, ":") {
			break // Nachbarschaftenliste erreicht
		}
		zellen := strings.Fields(line)
		kennwert, err := strconv.ParseInt(zellen[1], 10, 64)
		if err != nil {
			return nil, err
		}
		laenge, err := strconv.ParseFloat(zellen[2], 64)
		if err != nil {
			return nil, err
		}
		breite, err := strconv.ParseFloat(zellen[3], 64)
		if err != nil {
			return nil, err
		}
		staat := model.Staat{
			ID:       zellen[0],
			Kennwert: float64(kennwert),
			Position: model.Koordinate{
				X: laenge,
				Y: breite,
			},
		}
		container.Staaten = append(container.Staaten, staat)
	}

	container.Nachbarschaften = map[string]map[string]bool{}
	parseNachbarschaft(line, container.Nachbarschaften)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // Ignore comment
		}
		parseNachbarschaft(line, container.Nachbarschaften)
	}

	return &container, nil
}

func parseNachbarschaft(line string, n model.Nachbarschaften) {
	zeile := strings.Split(line, ":")
	n[zeile[0]] = map[string]bool{}
	nachbarn := strings.Fields(zeile[1])
	for _, nachbar := range nachbarn {
		n[zeile[0]][nachbar] = true
	}
}
