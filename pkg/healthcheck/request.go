package healthcheck

import (
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"sync"
	"time"
)

func SendRequests() {
	log.Println("About to send requests to check service availability")

	var wg sync.WaitGroup
	for _, service := range MainConfig.Services {
		wg.Add(1)
		go SendRequest(&wg, service)
	}

	wg.Wait()
	log.Println("Completed all service checks")
}

func SendContinuousRequests() {
	cronScheduler := cron.New()
	cronScheduler.AddFunc("* * * * *", SendRequests)
	cronScheduler.Start()
	time.Sleep(24 * time.Hour)
	cronScheduler.Stop()
}

func SendContinuousRequestsTickerVersion() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			log.Println("About to send requests to services")
			SendRequests()
		}
	}()

	// wait for 24 hours to completely finish
	time.Sleep(24 * time.Hour)
	ticker.Stop()
}

func SendRequest(wg *sync.WaitGroup, service Service) bool {
	defer wg.Done()
	log.Printf("Sending %s request to %s with url %s", service.Request.Method, service.Name, service.Url)

	var response *http.Response
	var err error
	if service.Request.Method == "GET" {
		response, err = http.Get(service.Url)
	} else if service.Request.Method == "HEAD" {
		response, err = http.Head(service.Url)
	}

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Got response %s from %s\n", response.Status, service.Name)
	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
