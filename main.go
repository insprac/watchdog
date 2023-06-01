package main

import (
	"github.com/insprac/watchdog/jobs"
	"github.com/robfig/cron"
)

func main() {
	runCronJobs()

	wait := make(chan struct{})
	<-wait
}

func runCronJobs() {
	c := cron.New()
	c.AddFunc("@every 1m", jobs.PriceWatchJob)
	c.AddFunc("@every 1m", jobs.SiloWatchJob)
	c.Start()
}
