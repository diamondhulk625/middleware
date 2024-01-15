package service

import "time"

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	t := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(t)

	return int64(dur.Hours()) / 24

}
