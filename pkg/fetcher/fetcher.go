package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Capstain/coinsmarketcup_fetcher/pkg/currency"
	"github.com/Capstain/coinsmarketcup_fetcher/pkg/token"
)

const (
	URL_TMPL       = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?slug=%s&convert=%s"
	HEADER_API_KEY = "X-CMC_PRO_API_KEY"
)

type Rate struct {
	Status json.RawMessage `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func FetchRate(t *token.Token, curr *currency.Currency) float64 {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(URL_TMPL, t.Code, curr.Code),
		nil,
	)
	if err != nil {
		log.Fatal("Fail on create request", err)
	}

	req.Header.Add(HEADER_API_KEY, os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Fail on fetch CMC response", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Fail on parsing response body")
		return 0
	}

	var rate Rate
	json.Unmarshal(body, &rate)

	re := regexp.MustCompile(`"price":\s*(\d*\.*\d*)`)
	q := re.Find(rate.Data)

	price, err := strconv.ParseFloat(strings.Split(string(q), ":")[1], 64)
	if err != nil {
		log.Fatal("Can't parse price. ", err)
		return 0
	}

	return price
}
