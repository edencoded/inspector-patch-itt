package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

var expRegex = regexp.MustCompile(`exploits/`)

type EDBScraper struct {
	BaseURL      string
	DataChannel  chan []EDBModel
	ErrorChannel chan error
	wg           sync.WaitGroup
	expMutex     sync.RWMutex
	exploits     []EDBModel
}

func newEBDScraper() EDBScraper {
	return EDBScraper{BaseURL: "https://www.exploit-db.com", DataChannel: make(chan []EDBModel), ErrorChannel: make(chan error)}
}

func (s *EDBScraper) getData(url string) {

	var urlNos []int
	urls := getUrls(url)
	logFile, err := os.ReadFile("Files/Logs/EBD_log.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range urls {
		urlNo, err := strconv.ParseInt(strings.Split(val, "/")[2], 10, 64)
		if err != nil {
			urlNo = 0
		}
		urlNos = append(urlNos, int(urlNo))
	}
	sort.Ints(urlNos)

	if strings.Trim(string(logFile), " ") == strings.Trim(fmt.Sprint(urlNos[len(urlNos)-1]), " ") {
		fmt.Println("Already in EBD Logs")
		return
	}
	s.wg.Add(len(urls))

	for _, url := range urls {
		go s.scrapeExploit(&s.wg, url)
	}

	s.wg.Wait()
	if len(s.exploits) == 0 {
		return
	} else {
		s.DataChannel <- s.exploits
	}

}

func (s *EDBScraper) saveData() {
	for {
		select {
		case data := <-s.DataChannel:
			var urlNos []int
			fo, err := os.Create("Files/EDB/edb-" + formatDateTime(time.Now()) + ".json")
			if err != nil {
				log.Fatal(err)
			} else {
				expJson, err := json.Marshal(map[string][]EDBModel{"exploits": data})
				if err != nil {
					log.Fatal(err)
				}
				fo.Write(expJson)
				for _, val := range data {
					urlNo, err := strconv.Atoi(strings.Split(val.Url, "/")[4])
					if err != nil {
						urlNo = 0
					}
					urlNos = append(urlNos, int(urlNo))
				}
				sort.Ints(urlNos)
				logFile, err := os.Create("Files/Logs/EBD_log.txt")
				if err != nil {
					log.Fatal(err)
				}
				logFile.Write([]byte(fmt.Sprint(urlNos[len(urlNos)-1])))
			}
		}
	}

}

func getUrls(domain string) []string {
	var expUrls []string
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	var bodyNodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Navigate(domain), chromedp.Sleep(2*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
		err := chromedp.Run(ctx, chromedp.Nodes("#exploits-table tbody a", &bodyNodes, chromedp.ByQueryAll))
		if err != nil {
			log.Fatal(err)
		}
		for _, node := range bodyNodes {
			link := node.AttributeValue("href")
			if expRegex.MatchString(strings.ToLower(link)) {
				expUrls = append(expUrls, link)
			}
		}
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}

	return expUrls
}

func (s *EDBScraper) scrapeExploit(wg *sync.WaitGroup, url string) {

	defer wg.Done()
	c := colly.NewCollector()
	c.OnHTML("h1.card-title.text-secondary.text-center", func(h *colly.HTMLElement) {
		expTitle := strings.TrimSpace(h.Text)
		exploit := EDBModel{Title: expTitle, Url: h.Request.URL.String()}
		s.expMutex.Lock()
		s.exploits = append(s.exploits, exploit)
		s.expMutex.Unlock()
	})

	c.Visit(s.BaseURL + url)

}
