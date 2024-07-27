package util

import "github.com/Solituderr/autoenv/util/cron"

var Cronjob cron.CronJobService

func Init() {
	Cronjob = cron.NewCronJobService()
}
