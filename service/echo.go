package service

import "go.uber.org/zap"

type EchoService struct {
	log *zap.Logger
}

func newEchoService(log *zap.Logger) *EchoService {
	return &EchoService{log: log}
}

func (s *EchoService) CallService() {
	s.log.Info("Call service")
}
