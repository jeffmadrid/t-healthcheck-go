package main

import (
	"github.com/jeffmadrid/healthcheck-one/pkg/healthcheck"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	healthcheck.ReadConfig()
	healthcheck.SendRequests()
	elapsedTime := time.Since(startTime)

	log.Printf("Completed Full Program Execution in %s\n", elapsedTime)
}
