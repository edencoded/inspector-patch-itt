[![Go Report Card](https://goreportcard.com/badge/github.com/swxft/inspector-patch-it)](https://goreportcard.com/report/github.com/swxft/inspector-patch-it)
# Inspector Patch It
>A Security Exploit Aggregator

This project aims to help development teams stay informed on the latest security exploits so that they can systematically become informed and delegate patch development work.
[Watch my Proposal for this program](https://www.loom.com/share/b043e635b55a4fa1b036230c8efa3c3d?sid=50163045-457f-43ff-8734-aa4847be2707)

[![Loom](loom-screenshot.png)](https://www.loom.com/share/b043e635b55a4fa1b036230c8efa3c3d?sid=50163045-457f-43ff-8734-aa4847be2707)
This prgram pulls in exploits from the [National Vulnerability Database (NVD) API](https://nvd.nist.gov/developers/vulnerabilities) and scrapes the ExploitsDB website for latest exploits . Check `vulns.json` for NVD API the results.
## How to Run 
clone this repo
```
git clone https://github.com/swxft/inspector-patch-it.git
```
```
go run .
```
`vulns.json` should then be created with the latest pull of exploits from 

## How to Run the Exploit DB Standalone Scraper

Install puppeteer with 
```
npm install puppeteer
```
We will then need to install `gocolly` and `chromedp` dependencies
```
go get github.com/chromedp/cdproto/cdp
```
```
go get github.com/chromedp/chromedp
```
run the `main.go` script
```
go run .
```
