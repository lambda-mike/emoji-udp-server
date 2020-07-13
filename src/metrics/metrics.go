package metrics

type MetricsService struct {
	valid   int
	invalid int
}

func (m *MetricsService) GetReport() string {
	return "TODO"
}

func (m *MetricsService) IncValid() {
}

func (m *MetricsService) IncInvalid() {
}

func Create() *MetricsService {
	return &MetricsService{}
}
