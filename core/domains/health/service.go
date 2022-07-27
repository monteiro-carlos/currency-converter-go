package health

import (
	"log"
	"math"
	"time"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/currency/repository"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health/enums"
	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health/models"
	"github.com/pkg/errors"
)

type ServiceI interface {
	CheckDatabase() models.Check
}

type Service struct {
	repository repository.ServiceI
}

const (
	okMsg, notOkMsg = "This service is OK", "This service is down"
)

func NewService(
	repository repository.ServiceI,
) (*Service, error) {
	if repository == nil {
		return nil, errors.New("repository is down")
	}

	return &Service{
		repository: repository,
	}, nil
}

func (h *Service) CheckDatabase() models.Check {
	startTime := time.Now()
	dbStatus, err := h.repository.DataBaseHealthCheck()
	endTime := time.Now()

	var status bool
	var msg string

	rawDuration := endTime.Sub(startTime).Seconds()
	duration := formatDuration(rawDuration)

	if dbStatus == "UP" {
		msg = okMsg
		status = true
	} else {
		status = false
		msg = notOkMsg
		err = errors.New("repository is down")
	}

	if err != nil {
		log.Print("repository is down")
	}

	return models.Check{
		Name:         "Database Check",
		Healthy:      status,
		Msg:          msg,
		Criticality:  enums.Criticality(1),
		ResponseTime: duration,
		CheckedAt:    time.Now(),
	}
}

func formatDuration(duration float64) float64 {
	const div = 1000000

	return math.Round(duration*div) / div
}
