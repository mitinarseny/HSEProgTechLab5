package report

import (
	"encoding/csv"
	"os"
)

type csvReporter struct {
	f *os.File
	w *csv.Writer
}

func newCsvReporter(filename string) (*csvReporter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(f)
	if err := w.Write([]string{"id", "timestamp"}); err != nil {
		return nil, err
	}
	return &csvReporter{
		f: f,
		w: w,
	}, nil
}

func (r *csvReporter) Report(e *Event) error {
	return r.w.Write(e.ToCSV())
}

func (r *csvReporter) Close() error {
	r.w.Flush()
	if err := r.w.Error(); err != nil {
		return err
	}
	return r.f.Close()
}
