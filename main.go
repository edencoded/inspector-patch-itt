package main

import "time"

func main() {

	ticker := time.NewTicker(12 * time.Hour)

	s1 := newNVDScraper()
	s2 := newEBDScraper()

	go s1.getData(s1.BaseURL)
	go s2.getData(s2.BaseURL)

	go s1.saveData()
	go s2.saveData()

	for {
		select {
		case <-ticker.C:
			go s1.getData(s1.BaseURL)
			s1.updateTime()
			go s2.getData(s2.BaseURL)

		}
	}

}
