package metrics

import (
	"fmt"
	"sync"
)

type MetricsService struct {
	valid   uint
	invalid uint
	mtx     sync.RWMutex
}

func (m *MetricsService) GetReport() string {
	m.mtx.RLock()
	template := "Valid inputs: %d\nInvalid inputs: %d"
	report := fmt.Sprintf(template, m.valid, m.invalid)
	m.mtx.RUnlock()
	return report
}

func (m *MetricsService) IncValid() {
	m.mtx.Lock()
	m.valid++
	m.mtx.Unlock()
}

func (m *MetricsService) IncInvalid() {
	m.mtx.Lock()
	m.invalid++
	m.mtx.Unlock()
}

func Create() *MetricsService {
	return &MetricsService{}
}
