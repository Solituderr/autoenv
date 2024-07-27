package cron

import (
	"errors"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func (c *CronJobImpl) SendByInterval(name string, interval time.Duration, f func()) error {
	if _, ok := c.mp.Load(name); ok {
		return errors.New("已存在该name的cronjob")
	} else {
		eid := c.cron.Schedule(cron.Every(interval), cron.FuncJob(f))
		c.mp.Store(name, eid)
		log.Printf("name: %s, 已加入cronjob", name)
	}
	return nil
}

func (c *CronJobImpl) Send(name string, spec string, f func()) error {
	if _, ok := c.mp.Load(name); ok {
		return errors.New("已存在该name的cronjob")
	} else {
		eid, err := c.cron.AddFunc(spec, f)
		if err != nil {
			return err
		}
		c.mp.Store(name, eid)
		log.Printf("name: %s, 已加入cronjob", name)
	}
	return nil
}

func (c *CronJobImpl) Close(name string) {
	if id, ok := c.mp.Load(name); ok {
		c.cron.Remove(id.(cron.EntryID))
		c.mp.Delete(name)
	}
}
