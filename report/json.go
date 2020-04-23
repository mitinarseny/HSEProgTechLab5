package report

import (
	"encoding/json"
	"os"
)

type jsonReporter struct {
	f *os.File
	encoder *json.Encoder
}

func newJsonReporter(filename string) (*jsonReporter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &jsonReporter{
		f: f,
		encoder: json.NewEncoder(f),
	}, nil
}

func (r *jsonReporter) Report(e *Event) error {
	return r.encoder.Encode(e)
}

func (r *jsonReporter) Close() error {
	return r.f.Close()
}
