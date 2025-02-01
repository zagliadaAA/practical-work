package service_provider

import (
	"medicalCenter/internal/pkg/timer"
)

func (sp *ServiceProvider) getTimer() *timer.Timer {
	if sp.timer == nil {
		sp.timer = timer.NewTimer()
	}

	return sp.timer
}
