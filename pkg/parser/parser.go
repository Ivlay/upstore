package htmlParser

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ivlay/upstore"
	"github.com/PuerkitoBio/goquery"
)

type HtmlParser struct {
	url string
}

func New(url string) *HtmlParser {
	return &HtmlParser{url: url}
}

func (p *HtmlParser) PrepareProduct() upstore.Product {
	var product upstore.Product

	doc, err := p.getMainDoc()
	if err != nil {
		log.Fatal(err)
	}

	pr := regexp.MustCompile("[0-9]+").FindString(strings.ReplaceAll(doc.Find(".product-price").Text(), " ", ""))

	price, err := strconv.ParseFloat(pr, 64)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find(".product-title").Text()

	product.Price = int(price)
	product.Title = title
	product.PriceId = "macbookair15_m3_10gpu_8-512gb_midnight"
	product.OldPrice = int(price)

	return product
}

func (p *HtmlParser) getMainDoc() (*goquery.Selection, error) {
	res, err := http.Get(p.url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Close connection")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	body := doc.Find("html body")

	return body, nil
}
