package models

import (
	"time"

	"github.com/monteiro-carlos/eng-gruposbf-backend-golang/core/domains/health/enums"
)

type Check struct {
	Name         string            `json:"name"`
	Healthy      bool              `json:"healthy"`
	Msg          string            `json:"msg"`
	Criticality  enums.Criticality `json:"criticality"`
	ResponseTime float64           `json:"responseTime,omitempty"`
	CheckedAt    time.Time         `json:"checkedAt"`
}

type HealthCheckReadiness struct {
	Checks    []Check   `json:"checks,omitempty"`
	Healthy   bool      `json:"healthy"`
	CheckedAt time.Time `json:"checkedAt"`
	Status    int       `json:"status"`
}

type HealthCheckLiveness struct {
	Healthy   bool      `json:"healthy"`
	CheckedAt time.Time `json:"checkedAt"`
	Status    int       `json:"status"`
}
