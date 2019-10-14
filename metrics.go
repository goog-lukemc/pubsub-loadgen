package main

import (
	"sync/atomic"
	"time"
)

type metrics struct {
	sent       uint64
	received   uint64
	start      time.Time
	psSent     uint64
	psReceived uint64
	roundTrip  uint64
	sendTimer  int64
}

func (m *metrics) average() {
	atomic.StoreUint64(&m.psSent, uint64(m.sent/uint64(time.Since(m.start).Seconds())))
	atomic.StoreUint64(&m.psReceived, uint64(m.received/uint64(time.Since(m.start).Seconds())))
}

func (m *metrics) addSent() {
	defer atomic.AddUint64(&m.sent, 1)
}

func (m *metrics) addReceived() {
	defer atomic.AddUint64(&m.received, 1)
}

func (m *metrics) updateRoundTrip(v uint64) {
	dt := uint64((uint64(time.Now().UnixNano()) - v) / uint64(time.Millisecond))
	defer atomic.StoreUint64(&dt, v)
}

func (m *metrics) updateSendTimer(v int64) {
	defer atomic.StoreInt64(&m.sendTimer, v)
}
