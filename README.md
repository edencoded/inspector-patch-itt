# Inspector Patch It
This prgram pulls in exploits from the [National Vulnerability Database (NVD) API](https://nvd.nist.gov/developers/vulnerabilities). Check `vulns.json` for NVD API the results.
## How to Run 
clone this repo
```
git clone https://github.com/swxft/inspector-patch-it.git
```
run the main script
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
go run main.go
```