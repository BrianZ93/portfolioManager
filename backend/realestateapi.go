package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

type RealEstatePost struct {
	Id          float64 	`json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Lien        float64 `json:"lien"`
	Rate        float64 `json:"rate"`
	Years       float64 `json:"years"`
	Value       float64 `json:"value"`
	MonthsLeft  float64 `json:"monthsleft"`
	Type 		string  `json:"type"`
}

type RealEstate []RealEstatePost

var CurrentProperties RealEstate
var CurrentProperty RealEstatePost
var TempProperties RealEstate

func checkREFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		logrus.Print("New RE portfolio detected, initializing...")
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
				Type: 		 "Empty",
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

	var redata interface{}

	json.Unmarshal(file, &redata)

	redataslice, ok := redata.([]interface{})

	if !ok {
		logrus.Print("cannot convert properties JSON objects")
		os.Exit(1)
	}

	logrus.Print("Number of JSON objects: ", len(redataslice))

	logrus.Print(redata, "initial RE data")

	var CurFloat float64
	var CurStr string

	var CurrentIndex = 0

	DataProceed := true

	for _, obj := range redataslice {
		objMap, ok := obj.(map[string]interface{})

		logrus.Print(objMap["id"])
		logrus.Print(reflect.TypeOf(objMap["id"]))

		if !ok {
			logrus.Print("cannot convert RE interface{} to type map[string]interface{}")
		}

		// The below if statements check on the data type to see if an upload is possible
		if res, ok := objMap["id"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Id = CurFloat
		}
		 

		if res, ok := objMap["description"].(string); !ok {
			CurStr = res
			DataProceed = false
			CurrentProperty.Description = CurStr
		}
		 

		if res, ok := objMap["price"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Price = CurFloat
		}

		 

		if res, ok := objMap["lien"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Lien = CurFloat
		}

		 

		if res, ok := objMap["rate"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Rate = CurFloat
		}
		 

		if res, ok := objMap["years"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Years = CurFloat
		}
		 

		if res, ok := objMap["value"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.Value = CurFloat
		}
		 

		if res, ok := objMap["monthsleft"].(float64); !ok {
			CurFloat = res
			DataProceed = false
			CurrentProperty.MonthsLeft = CurFloat
		}
	 

		if res, ok := objMap["type"].(string); !ok {
			CurStr = res
			DataProceed = false
			CurrentProperty.Type = CurStr
		}

		 

		if DataProceed {
			CurrentProperties = append(CurrentProperties, CurrentProperty)
		}

		logrus.Info(CurrentProperties)

		CurrentIndex++
	}

}

func RERetrieve(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	logrus.Print(CurrentProperties, "retrieve current properties")

	json.NewEncoder(w).Encode(CurrentProperties)
}

func addProperty(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var propertyData RealEstatePost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Property data at its memory location
	decoder.Decode(&propertyData)

	logrus.Print("++++++++++++++++++",propertyData)

	TempProperties = nil

	if CurrentProperties == nil {
		TempProperties = append(TempProperties, RealEstatePost{Id: 0, Description: propertyData.Description, Price: propertyData.Price, Lien: propertyData.Lien, Rate: propertyData.Rate, Years: propertyData.Years, Value: propertyData.Value, MonthsLeft: propertyData.MonthsLeft, Type: propertyData.Type})
		CurrentProperties = TempProperties

		file, _ := json.MarshalIndent(CurrentProperties, "", " ")

		_ = ioutil.WriteFile("Properties.json", file, 0644)
	} else {
		CurrentProperties = append(CurrentProperties, RealEstatePost{Id: float64(len(CurrentProperties)), Description: propertyData.Description, Price: propertyData.Price, Lien: propertyData.Lien, Rate: propertyData.Rate, Years: propertyData.Years, Value: propertyData.Value, MonthsLeft: propertyData.MonthsLeft, Type: propertyData.Type})

		file, _ := json.MarshalIndent(CurrentProperties, "", " ")

		_ = ioutil.WriteFile("Properties.json", file, 0644)
	}

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	logrus.Print(CurrentProperties, " current properties")

	json.NewEncoder(w).Encode(CurrentProperties)
}

func modifyProperty(w http.ResponseWriter, request *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Sets the appropriate data type for the JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var propertyData RealEstatePost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Property data at its memory location
	decoder.Decode(&propertyData)

	for property := range CurrentProperties {
		if propertyData.Id == CurrentProperties[property].Id {
			CurrentProperties[property].Id = propertyData.Id
			CurrentProperties[property].Description = propertyData.Description
			CurrentProperties[property].Price = propertyData.Price
			CurrentProperties[property].Lien = propertyData.Lien
			CurrentProperties[property].Rate = propertyData.Rate
			CurrentProperties[property].Years = propertyData.Years
			CurrentProperties[property].Value = propertyData.Value
			CurrentProperties[property].MonthsLeft = propertyData.MonthsLeft
			CurrentProperties[property].Type = propertyData.Type
		}
	}

	file, _ := json.MarshalIndent(CurrentProperties, "", " ")

	_ = ioutil.WriteFile("Properties.json", file, 0644)
}

func deleteProperty(w http.ResponseWriter, request *http.Request) {
	// Enabling Cors
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Initializing API variables
	var propertyData RealEstatePost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to Property data at its memory location
	decoder.Decode(&propertyData)

	if len(CurrentProperties) == 1 {
		TempProperties = nil
		CurrentProperties[0].Id = 0
		CurrentProperties[0].Description = ""
		CurrentProperties[0].Price = 0
		CurrentProperties[0].Lien = 0
		CurrentProperties[0].Rate = 0
		CurrentProperties[0].Years = 0
		CurrentProperties[0].Value = 0
		CurrentProperties[0].MonthsLeft = 0
		CurrentProperties[0].Type = "Empty"

	} else {
		TempProperties = nil
		for property := range CurrentProperties {
			if propertyData.Id == CurrentProperties[property].Id {
				continue
			} else {
				TempProperties = append(TempProperties, CurrentProperties[property])
			}
		}

		CurrentProperties = TempProperties
	}

	file, _ := json.MarshalIndent(CurrentProperties, "", " ")

	_ = ioutil.WriteFile("Properties.json", file, 0644)

}
