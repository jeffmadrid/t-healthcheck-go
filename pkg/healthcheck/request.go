package healthcheck

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func SendRequests() {
	var wg sync.WaitGroup
	for _, service := range MainConfig.Services {
		wg.Add(1)
		go SendRequest(&wg, service)
	}

	wg.Wait()
}

func SendContinuousRequests() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			log.Println("About to send requests to services")
			SendRequests()
			time.Sleep(1 * time.Minute)
		}
	}()

	// wait for 24 hours to completely finish
	time.Sleep(24 * time.Hour)
	ticker.Stop()
}

func SendRequest(wg *sync.WaitGroup, service Service) {
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
}
