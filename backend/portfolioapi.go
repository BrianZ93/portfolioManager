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

// Defining JSON Structures
type EquityPost struct {
	Ticker      string  `json:"ticker"`
	Shares      float64 `json:"shares"`
	Price       float64 `json:"price"`
	PriceLoaded bool    `json:"priceloaded"`
	Value       float64 `json:"value"`
	Message     string  `json:"message"`
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
		logrus.Print("cannot convert equities JSON objects")
		os.Exit(1)
	}

	var CurTicker string
	var TickerProceed = false
	var CurShares float64
	var SharesProceed = false
	var CurPrice float64
	var PriceProceed = false
	var CurValue float64
	var ValueProceed = false

	var CurrentIndex = 0

	for _, obj := range dataSlice {
		objMap, ok := obj.(map[string]interface{})

		if !ok {
			logrus.Print("cannot convert interface{} to type map[string]interface{}")
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

		CurrentIndex++

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

	logrus.Print(CurrentEquities, "retrieve current equities")

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

	TempEquities = nil

	if len(CurrentEquities) == 1 && CurrentEquities[0].Ticker == "0" && CurrentEquities[0].Shares == 0 {
		TempEquities = append(TempEquities, EquityPost{Ticker: equityData.Ticker, Shares: equityData.Shares, Price: equityData.Price})
		CurrentEquities = TempEquities

		updatePrices()

		file, _ := json.MarshalIndent(CurrentEquities, "", " ")

		_ = ioutil.WriteFile("Equities.json", file, 0644)

	} else {
		for i := 0; i < len(CurrentEquities); i++ {
			if equityData.Ticker == CurrentEquities[i].Ticker {
				if equityData.Shares > 0 {
					logrus.Print("Security Already Exists - Shares Added")
					if CurrentEquities[i].Shares+equityData.Shares >= 0 {
						CurrentEquities[i].Shares += equityData.Shares

						updatePrices()

						file, _ := json.MarshalIndent(CurrentEquities, "", " ")

						_ = ioutil.WriteFile("Equities.json", file, 0644)
					}
					break
				} else {
					logrus.Print("No Security Added - No Shares to add")
					break
				}

			} else if equityData.Shares == 0 || equityData.Ticker == "" {
				logrus.Print("Error Uploading Equity: Incomplete Data Sent")
			} else {
				if i == len(CurrentEquities)-1 {

					CurrentEquities = append(CurrentEquities, EquityPost{Ticker: equityData.Ticker, Shares: equityData.Shares, Price: equityData.Price})

					updatePrices()

					file, _ := json.MarshalIndent(CurrentEquities, "", " ")

					_ = ioutil.WriteFile("Equities.json", file, 0644)

					break
				} else {
					continue
				}

			}
		}
	}

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	logrus.Print(CurrentEquities, " current")

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

		updatePrices()

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
		TempEquities = nil
		CurrentEquities[0].Ticker = "0"
		CurrentEquities[0].Shares = 0
		CurrentEquities[0].Price = 0
		CurrentEquities[0].PriceLoaded = false
		CurrentEquities[0].Value = 0

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

	updatePrices()

	file, _ := json.MarshalIndent(CurrentEquities, "", " ")

	_ = ioutil.WriteFile("Equities.json", file, 0644)

}

func updatePrices() {

	for i := 0; i < len(CurrentEquities); i++ {
		q, err := quote.Get(CurrentEquities[i].Ticker)
		if err != nil {
			CurrentEquities[i].PriceLoaded = false
			CurrentEquities[i].Value = 0
		}

		CurrentEquities[i].PriceLoaded = true
		CurrentEquities[i].Value = q.RegularMarketPrice
	}
}