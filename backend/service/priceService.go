package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type price struct {
	client *http.Client
}

var (
	priceService *price
	priceOnce    = sync.OnceFunc(func() {
		priceService = new(price)
		priceService.client = &http.Client{}
	})
)

func GetPriceService() *price {
	priceOnce()
	return priceService
}

var (
	//TODO
	signalHeader map[string]map[string][]string = map[string]map[string][]string{"7d": map[string][]string{"Accepts": []string{"application/json"}, "X-CMC_PRO_API_KEY": []string{"79941117-fc00-4b38-b85f-96b29060a1ba"}}}
	singleValues map[string]url.Values          = map[string]url.Values{"7d": url.Values{}}
)

func (p *price) getpricesForHistorical(interval string) (any, error) {
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/historical", nil)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")
	q.Add("interval", "7d")
	req.URL.RawQuery = q.Encode()

	resp, err := p.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		return nil, err
	}
	io.ReadAll()

}
