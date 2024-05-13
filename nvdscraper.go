package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type NVDScraper struct {
	BaseURL      string
	PageSize     int
	NumPages     int
	CurrPage     int
	DataChannel  chan NVDResponse
	ErrorChannel chan error
	startDate    string
	currDate     string
}

func newNVDScraper() NVDScraper {
	currDT := time.Now()
	startDT := time.Now().AddDate(0, 0, -10)
	currTime := formatDateTime(currDT)
	startTime := formatDateTime(startDT)
	baseUrl := "https://services.nvd.nist.gov/rest/json/cves/2.0?pubStartDate=" + startTime + "&pubEndDate=" + currTime
	return NVDScraper{
		BaseURL:      baseUrl,
		currDate:     currTime,
		startDate:    startTime,
		DataChannel:  make(chan NVDResponse),
		ErrorChannel: make(chan error),
	}
}

func (s *NVDScraper) getData(url string) {

	var apiResponse NVDResponse
	resp, err := http.Get(s.BaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&apiResponse)

	if err != nil {
		s.ErrorChannel <- err
	}
	logFile, err := os.ReadFile("Files/Logs/NVD_log.txt")
	if err != nil {
		log.Fatal(err)
	}
	if strings.Trim(apiResponse.Vulnerabilities[0].Cve.ID, " ") == strings.Trim(string(logFile), " ") {
		fmt.Println("Vulnerability already present in NVD logs")
		return
	}
	s.DataChannel <- apiResponse

}

func (s *NVDScraper) saveData() {

	for {
		select {
		case data := <-s.DataChannel:
			fo, err := os.Create("Files/NVD/NVD" + s.currDate + ".json")
			if err != nil {
				log.Fatal(err)
			} else {
				vJson, err := json.Marshal(map[string]any{"vulnerabilities": data})
				if err != nil {
					log.Fatal(err)
				} else {
					fo.Write(vJson)
					logFile, err := os.Create("Files/Logs/NVD_log.txt")
					if err != nil {
						log.Fatal(err)
					} else {
						logFile.Write([]byte(data.Vulnerabilities[0].Cve.ID))
					}
				}
			}
		case err := <-s.ErrorChannel:
			log.Fatal(err)
		}
	}

}

func (s *NVDScraper) updateTime() {
	s.currDate = formatDateTime(time.Now())
	s.startDate = formatDateTime(time.Now().AddDate(0, 0, -1))
}
