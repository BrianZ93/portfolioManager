package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type RealEstatePost struct {
	Id          float64 `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Lien        float64 `json:"lien"`
	Rate        float64 `json:"rate"`
	Years       float64 `json:"years"`
	Value       float64 `json:"value"`
	MonthsLeft  float64 `json:"monthsleft"`
}

type RealEstate []RealEstatePost

var CurrentProperties RealEstate
var CurrentProperty RealEstatePost

// Function to create the Real Estate POST page
func realEstatePage(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS
	enableCors(&w)

	fmt.Fprintf(w, "POST API Endpoint")
}

func checkREFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("New RE portfolio detected, initializing...")
		_, err := os.Create(filename)

		StarterREData := &[]RealEstatePost{
			{
				Id:          1,
				Description: "Home",
				Price:       300000,
				Lien:        240000,
				Rate:        .05,
				Years:       30,
				MonthsLeft:  360,
				Value:       330000,
			},
		}
		file, _ := json.MarshalIndent(StarterREData, "", " ")

		_ = ioutil.WriteFile("Properties.json", file, 0644)

		if err != nil {
			return err
		}

	}
	return nil
}

func readREFile() {
	filename := "Properties.json"
	err := checkREFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	// Iterating through the JSON objects

	var REData interface{}

	json.Unmarshal(file, &REData)

	REdataSlice, ok := REData.([]interface{})

	if !ok {
		fmt.Println("cannot convert properties JSON objects")
		os.Exit(1)
	}

	fmt.Println("Number of JSON objects: ", len(REdataSlice))

	fmt.Println(REData, "initial RE data")

	var CurFloat float64
	var CurStr string

	var CurrentIndex = 0

	DataProceed := true

	fmt.Println(DataProceed, "data proceed")

	for _, obj := range REdataSlice {
		objMap, ok := obj.(map[string]interface{})

		if !ok {
			fmt.Println("cannot convert RE interface{} to type map[string]interface{}")
		}

		// The below if statements check on the data tyrpe to see if an upload is possible
		if res, ok := objMap["Id"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Id = CurFloat
		}

		if res, ok := objMap["Description"].(string); !ok {
			CurStr = res
			DataProceed = false
			CurrentProperty.Description = CurStr
		}

		if res, ok := objMap["Price"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Price = CurFloat
		}

		if res, ok := objMap["Lien"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Lien = CurFloat
		}

		if res, ok := objMap["Rate"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Rate = CurFloat
		}

		if res, ok := objMap["Years"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Years = CurFloat
		}

		if res, ok := objMap["Value"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Value = CurFloat
		}

		if res, ok := objMap["MonthsLeft"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.MonthsLeft = CurFloat
		}

		if DataProceed {
			CurrentProperties = append(CurrentProperties, CurrentProperty)
		}

		CurrentIndex++
	}

}

func RERetrieve(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var REData RealEstatePost

	// Initializing the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Equity data at its memory location
	decoder.Decode(&REData)

	file, _ := json.MarshalIndent(CurrentProperties, "", " ")

	_ = ioutil.WriteFile("Properties.json", file, 0644)
}
