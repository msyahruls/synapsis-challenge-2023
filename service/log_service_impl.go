package service

import (
	"time"

	"synapsis-challange/model/domain"
	"synapsis-challange/repository"
)

type LogServiceImpl struct {
	LogRepository repository.LogRepository
}

type LogService interface {
	Create(actor, action string)
	FindAll(page, limit string) []domain.Logs
}

func NewLogService(logRepository repository.LogRepository) LogService {
	return &LogServiceImpl{
		LogRepository: logRepository,
	}
}

func (service *LogServiceImpl) Create(actor, action string) {

	dataLog := domain.Logs{
		Actor:     actor,
		Action:    action,
		Timestamp: time.Now(),
	}
	service.LogRepository.Create(dataLog)
}

func (service *LogServiceImpl) FindAll(page, limit string) []domain.Logs {

	log := service.LogRepository.FindAll(page, limit)

	return log
}
