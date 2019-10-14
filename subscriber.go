package main

import (
	"strconv"
	"time"
)

func subscriber() {
	ticker := time.NewTicker(time.Second)
	for pm := range inbound {
		select {
		case <-ticker.C:
			// Calculate the round trip time
			if s, ok := pm.Attributes["timestamp"]; ok {
				rt, _ := strconv.ParseUint(s, 10, 64)
				numbers.updateRoundTrip(rt)
			}
			//Add receive stats
			numbers.addReceived()

		default:

		}

	}

}
