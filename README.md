[![Go Report Card](https://goreportcard.com/badge/github.com/swxft/inspector-patch-itt)](https://goreportcard.com/report/github.com/swxft/inspector-patch-itt)
# Inspector Patch It
>A Security Exploit Aggregator

This project aims to help development teams stay informed on the latest security exploits so that they can systematically become informed and delegate patch development work.
[Watch my Proposal for this program](https://www.loom.com/share/b043e635b55a4fa1b036230c8efa3c3d?sid=50163045-457f-43ff-8734-aa4847be2707)

[![Loom](loom-screenshot.png)](https://www.loom.com/share/b043e635b55a4fa1b036230c8efa3c3d?sid=50163045-457f-43ff-8734-aa4847be2707)
This prgram pulls in exploits from the [National Vulnerability Database (NVD) API](https://nvd.nist.gov/developers/vulnerabilities) and scrapes the ExploitsDB website for latest exploits . Check the `/Files` directory for exploits.
## How to Run 
clone this repo
```
git clone https://github.com/swxft/inspector-patch-it.git
```
```
go run .
```
Check the `/Files` directory for exploits. This program works as long as it is running. Every 12 hours it is programmed to concurrenty check the 2 resources for updates. 
### Requirements

- [x] Builds, installs, and executes successfully
- [ ] B or higher on Go Report Card
- [x] Incorporates an external API or package
- [x] Persists data in a file or databse
- [x] README contains description
- [x] README contains screenshot or install instructions
- [x] README contains example of how to use this program
- [ ] 2 or more table-driven tests
- [ ] 1 or more benchmark tests
- [ ] All tests pass
---
## How to Run the Exploit DB Standalone Scraper
>(This would be a sidequest, not necessary for running the actual application.)
Navigate to the `scraper` branch
```
git checkout scraper 
```

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