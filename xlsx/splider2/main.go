package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() MyPageProcesser {
	return MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()

	name := query.Find(".lemmaTitleH1").Text()
	name = strings.Trim(name, " \t\n")

	summary := query.Find(".card-summary-content .para").Text()
	summary = strings.Trim(summary, " \t\n")

	// the entity we want to save by Pipeline
	p.AddField("name", name)
	p.AddField("summary", summary)
}

func (this MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

func main() {
	// spider input:
	//  PageProcesser ;
	//  task name used in Pipeline for record;
	spider.NewSpider(NewMyPageProcesser(), "TaskName").
		AddUrl("https://curve.fi/usecrv", "html").  // Start url, html is the responce type ("html" or "json")
		AddPipeline(pipeline.NewPipelineConsole()). // Print result on screen
		SetThreadnum(3).                            // Crawl request by three Coroutines
		Run()
	//sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
	// GetWithParams Params:
	//  1. Url.
	//  2. Responce type is "html" or "json" or "jsonp" or "text".
	//  3. The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
	//  4. The method is POST or GET.
	//  5. The postdata is body string sent to sever.
	//  6. The header is header for http request.
	//  7. Cookies
	//req := request.NewRequest("https://curve.fi/usecrv", "html", "", "GET", "", nil, nil, nil, nil)
	//pageItems := sp.GetByRequest(req)
	//pageItems := sp.Get("https://curve.fi/usecrv", "html")
	//fmt.Println(pageItems.)
	//bys, _ := json.Marshal(pageItems)
	//fmt.Println(string(bys))
}
