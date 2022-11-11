package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type DebtPost struct {
	Id     float64 `json:"id"`
	Title  string  `json:"title"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
	Term   float64 `json:"term"`
	Date   string  `json:"date"`
}

type Debt []DebtPost

var CurrentDebts Debt
var TemtDebts Debt

func CheckDebtFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		logrus.Info("New portfolio detected, initializing...")
		_, err := os.Create(filename)

		StarterDebtData := &[]DebtPost{
			{
				Id:     1,
				Title:  "Home",
				Type:   "Installment",
				Amount: 208000,
				Rate:   .02615,
				Term:   30,
				Date:   "10/01/2022",
			},
		}
		file, _ := json.MarshalIndent(StarterDebtData, "", " ")

		_ = ioutil.WriteFile("Debts.json", file, 0644)

		if err != nil {
			return err
		}

	}
	return nil
}

func readDebtFile() {

	filename := "Debts.json"
	err := CheckDebtFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	// Iterating through the JSON objects

	var debtdata interface{}

	json.Unmarshal(file, &debtdata)

	debtdataslice, ok := debtdata.([]interface{})

	if !ok {
		logrus.Error("cannot convert Debts JSON objects")
	}

	var CurId float64
	IdProceed := false
	var CurTitle string
	TitleProceed := false
	var CurType string
	TypeProceed := false
	var CurAmount float64
	AmountProceed := false
	var CurRate float64
	RateProceed := false
	var CurTerm float64
	TermProceed := false
	var CurDate string
	DateProceed := false

	var CurrentIndex = 0

	for _, obj := range debtdataslice {
		objMap, ok := obj.(map[string]interface{})

		if !ok {
			logrus.Error("cannot convert RE interface{} to type map[string]interface{}")
		}

		// The below if statements check on the data type to see if an upload is possible
		if res, ok := objMap["id"].(float64); ok {
			CurId = res
			IdProceed = true
		}

		if res, ok := objMap["title"].(string); ok {
			CurTitle = res
			TitleProceed = true
		}

		if res, ok := objMap["type"].(string); ok {
			CurType = res
			TypeProceed = true
		}

		if res, ok := objMap["amount"].(float64); ok {
			CurAmount = res
			AmountProceed = true
		}

		if res, ok := objMap["rate"].(float64); ok {
			CurRate = res
			RateProceed = true
		}

		if res, ok := objMap["term"].(float64); ok {
			CurTerm = res
			TermProceed = true
		}

		if res, ok := objMap["date"].(string); ok {
			CurDate = res
			DateProceed = true
		}

		if IdProceed && TitleProceed && TypeProceed && AmountProceed && RateProceed && TermProceed && DateProceed {
			CurrentDebts = append(CurrentDebts, DebtPost{Id: CurId, Title: CurTitle, Type: CurType, Amount: CurAmount, Rate: CurRate, Term: CurTerm, Date: CurDate})
		}

		CurrentIndex++
	}

}

func DebtRetrieve(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(CurrentDebts)
}

func addDebt(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var debtData DebtPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Debt data at its memory location
	decoder.Decode(&debtData)

	TemtDebts = nil

	if CurrentDebts == nil {
		TemtDebts = append(TemtDebts, DebtPost{Id: 0, Title: debtData.Title, Type: debtData.Type, Amount: debtData.Amount, Rate: debtData.Rate, Term: debtData.Term, Date: debtData.Date})
		CurrentDebts = TemtDebts

		file, _ := json.MarshalIndent(CurrentDebts, "", " ")

		_ = ioutil.WriteFile("Debts.json", file, 0644)
	} else {

		NewId := float64(len(CurrentDebts))

		IdAvailable := false
		IdInBatch := false

		for !IdAvailable {
			IdInBatch = false
			for _, debt := range CurrentDebts {
				if NewId == debt.Id {
					IdInBatch = true
				}
			}

			if !IdInBatch {
				IdAvailable = true
			} else {
				NewId++
			}
		}

		CurrentDebts = append(CurrentDebts, DebtPost{Id: NewId, Title: debtData.Title, Type: debtData.Type, Amount: debtData.Amount, Rate: debtData.Rate, Term: debtData.Term, Date: debtData.Date})

		file, _ := json.MarshalIndent(CurrentDebts, "", " ")

		_ = ioutil.WriteFile("Debts.json", file, 0644)
	}

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(CurrentDebts)
}

func modifyDebt(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for the JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var debtData DebtPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Debt data at its memory location
	decoder.Decode(&debtData)

	for debt := range CurrentDebts {
		if debtData.Id == CurrentDebts[debt].Id {
			CurrentDebts[debt].Id = debtData.Id
			CurrentDebts[debt].Title = debtData.Title
			CurrentDebts[debt].Type = debtData.Type
			CurrentDebts[debt].Amount = debtData.Amount
			CurrentDebts[debt].Rate = debtData.Rate
			CurrentDebts[debt].Term = debtData.Term
			CurrentDebts[debt].Date = debtData.Date
		}
	}

	file, _ := json.MarshalIndent(CurrentDebts, "", " ")

	_ = ioutil.WriteFile("Debts.json", file, 0644)
}

func deleteDebt(w http.ResponseWriter, request *http.Request) {
	// Enabling Cors
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var debtData DebtPost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Debt data at its memory location
	decoder.Decode(&debtData)

	if len(CurrentDebts) == 1 {
		TemtDebts = nil
		CurrentDebts[0].Id = 0
		CurrentDebts[0].Title = ""
		CurrentDebts[0].Type = "Empty"
		CurrentDebts[0].Amount = 0
		CurrentDebts[0].Rate = 0
		CurrentDebts[0].Term = 0
		CurrentDebts[0].Date = "Empty"

	} else {
		TemtDebts = nil
		for debt := range CurrentDebts {
			if debtData.Id == CurrentDebts[debt].Id {
				continue
			} else {
				TemtDebts = append(TemtDebts, CurrentDebts[debt])
			}
		}

		CurrentDebts = TemtDebts
	}

	file, _ := json.MarshalIndent(CurrentDebts, "", " ")

	_ = ioutil.WriteFile("Debts.json", file, 0644)

}
