package fetcher

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"challenge/models"
)

type apiResponse struct {
	Items    []models.StockAPIItem `json:"items"`
	NextPage string                `json:"next_page"`
}

func parseDollar(s string) float64 {
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func FetchAllStocks() ([]models.Stock, error) {
	var allStocks []models.Stock
	page := ""

	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_TOKEN")
	if apiURL == "" || apiKey == "" {
		return nil, errors.New("API_URL o API_KEY no están definidos en las variables de entorno")
	}

	for {
		url := apiURL
		if page != "" {
			url += "?next_page=" + page
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New("error al obtener datos de la API")
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var data apiResponse
		json.Unmarshal(body, &data)

		for _, item := range data.Items {
			stock := models.Stock{
				Ticker:     item.Ticker,
				Company:    item.Company,
				Brokerage:  item.Brokerage,
				Action:     item.Action,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				TargetFrom: parseDollar(item.TargetFrom),
				TargetTo:   parseDollar(item.TargetTo),
				Time:       item.Time,
			}
			allStocks = append(allStocks, stock)
		}

		if data.NextPage == "" {
			break
		}
		page = data.NextPage
	}

	return allStocks, nil
}
