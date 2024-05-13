package main

type Scraper interface {
	getData(string)
	saveData()
}
