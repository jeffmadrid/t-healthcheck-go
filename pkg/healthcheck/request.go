package healthcheck

import (
	"log"
	"net/http"
	"sync"
)

func SendRequests() {
	var wg sync.WaitGroup
	for _, service := range MainConfig.Services {
		wg.Add(1)
		go SendRequest(&wg, service)
	}

	wg.Wait()
}

func SendRequest(wg *sync.WaitGroup, service Service) {
	defer wg.Done()
	log.Printf("Sending request to %s with url %s", service.Name, service.Url)

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
