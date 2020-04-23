package report

import (
	"fmt"
	"strconv"
	"time"
)

type Event struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *Event) ToCSV() []string {
	return []string{
		strconv.Itoa(e.ID),
		e.Timestamp.String(),
	}
}

type Reporter interface {
	Report(*Event) error
	Close() error
}

type ReporterType int

const (
	JSONReporter = ReporterType(iota)
	CSVReporter
)

const (
	defaultJSONFilename = "report.json"
	defaultCSVFilename  = "report.csv"
)

func NewReporter(typ ReporterType) (Reporter, error) {
	switch typ {
	case JSONReporter:
		return newJsonReporter(defaultJSONFilename)
	case CSVReporter:
		return newCsvReporter(defaultCSVFilename)
	default:
		panic(fmt.Sprintf("unknown reporter type: %v", typ))
	}
}
