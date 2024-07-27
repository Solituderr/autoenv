package cron

import (
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type CronJobImpl struct {
	ch   chan func()
	mp   sync.Map
	cron *cron.Cron
}

// id must be unique
type CronJobService interface {
	SendByInterval(name string, interval time.Duration, f func()) error
	Send(name string, spec string, f func()) error
	Close(name string)
}

func NewCronJobService() CronJobService {
	c := cron.New()
	c.Start()
	return &CronJobImpl{
		ch:   make(chan func()),
		mp:   sync.Map{},
		cron: c,
	}
}
