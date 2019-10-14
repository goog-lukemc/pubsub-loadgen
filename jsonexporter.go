package main

import (
	"encoding/json"

	"go.opencensus.io/stats/view"
)

var (
	osEvents chan []byte
	osErrors chan error
	jexp     = JSONExporter{}
)

// JSONExporter
type JSONExporter struct{}

func newJSONExporter() *JSONExporter {
	if osEvents == nil {
		osEvents = make(chan []byte, 1000)
	}

	if osErrors == nil {
		osErrors = make(chan error, 1000)
	}

	return &jexp
}

// ExportView logs the view data.
func (e JSONExporter) ExportView(vd *view.Data) {
	bts, err := json.Marshal(vd)
	if err != nil {
		osErrors <- err
	}
	osEvents <- bts
}
