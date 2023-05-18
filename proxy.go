package binance

import (
	"sync"
	"time"
)

type Proxy struct {
	weightBalance int
	m             sync.Mutex
	ticker        *time.Ticker
}

var defaultProxy *Proxy

func NewProxy() *Proxy {
	ticker := time.NewTicker(1 * time.Minute)
	p := &Proxy{
		ticker:        ticker,
		weightBalance: 1200,
	}
	go func() {
		select {
		case <-ticker.C:
			p.m.Lock()
			p.weightBalance = 1200
			p.m.Unlock()
		}
	}()
	return p
}
func (p *Proxy) waitForWeight(weight int) {
	for {
		p.m.Lock()
		if p.weightBalance > weight {
			p.weightBalance -= weight
			p.m.Unlock()
			break
		}
		p.m.Unlock()
		time.Sleep(time.Second)
	}
}
