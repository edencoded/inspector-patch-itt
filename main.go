package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var nvdUrl = "https://services.nvd.nist.gov/rest/json/cves/2.0?pubStartDate=2024-05-03T00:00:00.000&pubEndDate=2024-05-10T00:00:00.000"
var startIndex = 0
var data []NVDResponse
var wg sync.WaitGroup
var dMutex sync.RWMutex

func main() {

	var response NVDResponse

	resp, err := http.Get(nvdUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	data = append(data, response)
	startIndex += response.ResultsPerPage
	wg.Add(response.TotalResults/response.ResultsPerPage - 1)
	for response.ResultsPerPage < response.TotalResults {
		url := nvdUrl + "&startIndex=" + strconv.Itoa(startIndex)
		startIndex += response.ResultsPerPage
		go getNVDData(&wg, url)
	}

	wg.Wait()

	fo, err := os.Create("vulns.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("saving file now")

	vJson, err := json.Marshal(map[string]any{"vulnerabilities": data})
	if err != nil {
		log.Fatal(err)
	}
	fo.Write(vJson)
	fmt.Println("Finished and file saved as vulns.json")
}

func getNVDData(wg *sync.WaitGroup, url string) {

	var apiResp NVDResponse

	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal((err))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatal(err)
	}
	dMutex.Lock()
	data = append(data, apiResp)
	dMutex.Unlock()
}
