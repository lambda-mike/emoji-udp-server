package metrics

import (
	"fmt"
	"log"
	"sync"
)

type MetricsService struct {
	valid   uint
	invalid uint
	mtx     sync.RWMutex
}

func (m *MetricsService) GetReport() string {
	log.Println("INFO GetReport")
	m.mtx.RLock()
	template := "Valid inputs: %d\nInvalid inputs: %d"
	report := fmt.Sprintf(template, m.valid, m.invalid)
	m.mtx.RUnlock()
	return report
}

func (m *MetricsService) IncValid() {
	log.Println("INFO IncValid")
	m.mtx.Lock()
	m.valid++
	m.mtx.Unlock()
}

func (m *MetricsService) IncInvalid() {
	log.Println("INFO IncInvalid")
	m.mtx.Lock()
	m.invalid++
	m.mtx.Unlock()
}

func Create() *MetricsService {
	log.Println("INFO MetricsService Create")
	return &MetricsService{}
}
