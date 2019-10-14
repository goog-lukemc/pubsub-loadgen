package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	CLEAR            = "\033[2J"
	SPACE            = " "
	BOARDER          = "-------------------------------------------------------\n"
	TOPICLABEL       = "| Topic: %s\n"
	MESSAGESSENT     = "|\tTotal Messages Sent\t->\t%v\n"
	ROUNDTRIPLATENCY = "|\tAverage Round Trip\t->\t%.1fms\n"
	SENTBYTESPERRPC  = "|\tBytes Per RPC\t\t->\t%.0fB\n"
	RTL              = "grpc.io/client/roundtrip_latency"
	SBR              = "grpc.io/client/sent_bytes_per_rpc"
	PMT              = "cloud.google.com/go/pubsub/published_messages"
	TOPIC            = "topic"
)

var (
	reportChan chan string
	ledger     sync.Map
	ot         sync.Once
	outs       []string = []string{TOPIC, RTL, SBR, PMT}
)

func writeEvents() {
	go func() {
		ot.Do(func() {
			printStatus(time.NewTicker(time.Second * 2))
		})
	}()
	for e := range vi {
		switch e.name {
		case RTL:
			ledger.Store(RTL, fmt.Sprintf(ROUNDTRIPLATENCY, e.distributionData.Mean))
		case PMT:
			ledger.Store(PMT, fmt.Sprintf(MESSAGESSENT, e.sumData.Value))
		case SBR:
			ledger.Store(SBR, fmt.Sprintf(SENTBYTESPERRPC, e.distributionData.Mean))
		}
		for k, v := range e.tags {
			if k == TOPIC {
				ledger.Store(k, fmt.Sprintf(TOPICLABEL, v))
			}

		}
	}
}

func printStatus(print *time.Ticker) {
	sb := strings.Builder{}

	for range print.C {
		sb.WriteString(CLEAR)
		sb.WriteString(BOARDER)
		for _, m := range outs {
			getFromLedger(m, &sb)
		}
		sb.WriteString(BOARDER)
		fmt.Printf("%s", sb.String())
	}
}

func getFromLedger(k string, sb *strings.Builder) {
	if v, ok := ledger.Load(k); ok {
		sb.WriteString(v.(string))
	}
}
