package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/piquette/finance-go/iter"
	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
)

// TODO - remove the json tags and see if you no longer need to use the interface hack to get the json to unmarshal correctly

// Defining JSON Structures
type EquityPost struct {
	Ticker      string  `json:"ticker"`
	Shares      float64 `json:"shares"`
	Price       float64 `json:"price"`
	PriceLoaded bool    `json:"priceloaded"`
	Value       float64 `json:"value"`
	Margin      float64 `json:"margin"`
}

type Equities []EquityPost

type Iter struct {
	*iter.Iter
}

var CurrentEquities Equities
var TempEquities Equities

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		logrus.Print("New portfolio detected, initializing...")
		_, err := os.Create(filename)

		StarterData := &[]EquityPost{
			{
				Ticker:      "AAPL",
				Shares:      100,
				Price:       0,
				PriceLoaded: false,
				Value:       0,
			},
			{
				Ticker:      "QQQ",
				Shares:      200,
				Price:       0,
				PriceLoaded: false,
				Value:       0,
			},
		}
		file, _ := json.MarshalIndent(StarterData, "", " ")

		_ = ioutil.WriteFile("Equities.json", file, 0644)

		if err != nil {
			return err
		}
	}
	return nil
}

func readFile() {

	filename := "Equities.json"
	err := checkFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	// Iterating through the JSON objects

	var data interface{}

	json.Unmarshal(file, &data)

	dataSlice, ok := data.([]interface{})

	if !ok {
		logrus.Info("cannot convert equities JSON objects")
	}

	var CurTicker string
	var TickerProceed = false
	var CurShares float64
	var SharesProceed = false
	var CurPrice float64
	var PriceProceed = false
	var CurValue float64
	var ValueProceed = false

	for _, obj := range dataSlice {
		objMap, ok := obj.(map[string]interface{})

		if !ok {
			logrus.Info("cannot convert interface{} to type map[string]interface{}")
		}

		if res, ok := objMap["ticker"].(string); ok {
			CurTicker = res
			TickerProceed = true
		}

		if res, ok := objMap["shares"].(float64); ok {
			CurShares = res
			SharesProceed = true
		}

		if res, ok := objMap["price"].(float64); ok {
			CurPrice = res
			PriceProceed = true
		}

		if res, ok := objMap["value"].(float64); ok {
			CurValue = res
			ValueProceed = true
		}

		if TickerProceed && SharesProceed && PriceProceed && ValueProceed {
			CurrentEquities = append(CurrentEquities, EquityPost{Ticker: CurTicker, Shares: CurShares, Price: CurPrice, Value: CurValue})

			// API Call to find equity values
		}

	}

}

func retrieveEquities(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(CurrentEquities)
}

func addEquity(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")

	// Initiatlizing API variables
	var equityData EquityPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Equity data at its memory location
	decoder.Decode(&equityData)

	q, err := quote.Get(equityData.Ticker)
	if err != nil {
		logrus.Error(q, "is not a recognized security")
	}

	TempEquities = nil

	if len(CurrentEquities) == 0 {
		TempEquities = append(TempEquities, EquityPost{Ticker: equityData.Ticker, Shares: equityData.Shares, Price: equityData.Price})
		CurrentEquities = TempEquities

		file, _ := json.MarshalIndent(CurrentEquities, "", " ")

		_ = ioutil.WriteFile("Equities.json", file, 0644)

	} else {
		for i := 0; i < len(CurrentEquities); i++ {
			if equityData.Ticker == CurrentEquities[i].Ticker {
				if equityData.Shares > 0 {
					logrus.Info("Security Already Exists - Shares Added")
					if CurrentEquities[i].Shares+equityData.Shares >= 0 {
						CurrentEquities[i].Shares += equityData.Shares
						CurrentEquities[i].Price = equityData.Price

						file, _ := json.MarshalIndent(CurrentEquities, "", " ")

						_ = ioutil.WriteFile("Equities.json", file, 0644)
					}
					break
				} else {
					logrus.Error("No Security Added - No Shares to add")
					break
				}

			} else if equityData.Shares == 0 || equityData.Ticker == "" {
				logrus.Error("Error Uploading Equity: Incomplete Data Sent")
			} else {
				if i == len(CurrentEquities)-1 {

					CurrentEquities = append(CurrentEquities, EquityPost{Ticker: equityData.Ticker, Shares: equityData.Shares, Price: equityData.Price})

					file, _ := json.MarshalIndent(CurrentEquities, "", " ")

					_ = ioutil.WriteFile("Equities.json", file, 0644)

					break
				} else {
					continue
				}

			}
		}
	}

	updatePrices()

	// TODO - When Voya's website is working look through PDF's to potential use to scrape through for updated portfolio Data

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(CurrentEquities)
}

func modifyEquity(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var equityData EquityPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Equity data at its memory location
	decoder.Decode(&equityData)

	for equity := range CurrentEquities {
		if equityData.Ticker == CurrentEquities[equity].Ticker {
			CurrentEquities[equity].Shares = equityData.Shares
			CurrentEquities[equity].Price = equityData.Price
		}

		file, _ := json.MarshalIndent(CurrentEquities, "", " ")

		_ = ioutil.WriteFile("Equities.json", file, 0644)
	}

}

func deleteEquity(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var equityData EquityPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Equity data at its memory location
	decoder.Decode(&equityData)

	if len(CurrentEquities) == 1 {

		CurrentEquities = nil

	} else {
		TempEquities = nil
		for equity := range CurrentEquities {
			if equityData.Ticker == CurrentEquities[equity].Ticker {
				continue
			} else {
				TempEquities = append(TempEquities, CurrentEquities[equity])
			}

		}

		CurrentEquities = TempEquities

	}

	file, _ := json.MarshalIndent(CurrentEquities, "", " ")

	_ = ioutil.WriteFile("Equities.json", file, 0644)

}

func updatePrices() {

	for i := 0; i < len(CurrentEquities); i++ {
		q, err := quote.Get(CurrentEquities[i].Ticker)
		if err != nil {
			CurrentEquities[i].PriceLoaded = false
			CurrentEquities[i].Value = 0
			logrus.Warn("Ticker not found")
		}

		CurrentEquities[i].PriceLoaded = true

		if q != nil {
			CurrentEquities[i].Value = q.RegularMarketPrice

		}

	}
}
