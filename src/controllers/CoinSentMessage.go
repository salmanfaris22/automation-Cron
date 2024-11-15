package controllers

import (
	"github.com/robfig/cron/v3"

	"main/src/service"
)

func RunCron() {
	c := cron.New()
	c.AddFunc("@every 00h00m2s", service.SentMessage)
	c.Start()

}
