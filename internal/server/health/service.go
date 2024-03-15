package health

type Service struct{}

func (s *Service) Get() *Health {
	return &Health{
		Status: "OK",
	}
}
