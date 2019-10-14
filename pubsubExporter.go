package main

import (
	"sync"
	"time"

	"go.opencensus.io/stats/view"
)

var (
	vi   chan viewInfo
	pexp = PubSubExporter{}
	once = sync.Once{}
)

// JSONExporter
type PubSubExporter struct{}

type viewInfo struct {
	name             string
	description      string
	start            time.Time
	end              time.Time
	distributionData *view.DistributionData
	countData        *view.CountData
	sumData          *view.SumData
	lastValueData    *view.LastValueData
	tags             map[string]string
}

func newPubSubExporter() *PubSubExporter {
	once.Do(newPVEXP)
	return &pexp
}

func newPVEXP() {
	vi = make(chan viewInfo, 1)

}

// ExportView logs the view data.
func (e PubSubExporter) ExportView(vd *view.Data) {
	cvi := viewInfo{}
	cvi.name = vd.View.Name
	cvi.description = vd.View.Description
	cvi.start = vd.Start
	cvi.end = vd.End
	cvi.tags = make(map[string]string)

	for _, row := range vd.Rows {
		switch v := row.Data.(type) {
		case *view.DistributionData:
			cvi.distributionData = v
		case *view.CountData:
			cvi.countData = v
		case *view.SumData:
			cvi.sumData = v
		case *view.LastValueData:
			cvi.lastValueData = v
		}
		for _, tag := range row.Tags {
			cvi.tags[tag.Key.Name()] = tag.Value
		}
	}
	vi <- cvi
}
