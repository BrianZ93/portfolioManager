package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type Tenant struct {
	Name          string    `json:"name"`
	LeaseStart    string    `json:"leasestart"`
	LeaseEnd      string    `json:"leaseend"`
	Expenses      []Expense `json:"expenses"`
	Revenues      []Revenue `json:"revenues"`
	Unit          string    `json:"unit"`
	CurrentTenant bool      `json:"currenttenant"`
	Id            float64   `json:"id"`
	SubId         float64   `json:"subid"`
}

type Expense struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

type Revenue struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
}

type RealEstatePost struct {
	Id               float64   `json:"id"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	Lien             float64   `json:"lien"`
	Rate             float64   `json:"rate"`
	Years            float64   `json:"years"`
	Value            float64   `json:"value"`
	MonthsLeft       float64   `json:"monthsleft"`
	Type             string    `json:"type"`
	Ownership        float64   `json:"ownership"`
	Date             string    `json:"date"`
	Tenants          []Tenant  `json:"tenants"`
	BuildingExpenses []Expense `json:"buildingexpenses"`
	BuildingRevenues []Revenue `json:"buildingrevenues"`
}

type RealEstate []RealEstatePost

var CurrentProperties RealEstate
var TempProperties RealEstate

func checkREFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		logrus.Info("New RE portfolio detected, initializing...")
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
				Type:        "Empty",
				Ownership:   100,
				Date:        "10/01/2022",
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
		logrus.Error("cannot convert properties JSON objects")
	}

	var CurId float64
	IdProceed := false
	var CurDesc string
	DescProceed := false
	var CurPrice float64
	PriceProceed := false
	var CurLien float64
	LienProceed := false
	var CurRate float64
	RateProceed := false
	var CurYears float64
	YearsProceed := false
	var CurValue float64
	ValueProceed := false
	var CurMonthsLeft float64
	MonthsLeftProceed := false
	var CurType string
	TypeProceed := false
	var CurOwnership float64
	OwnershipProceed := false
	var CurDate string
	DateProceed := false

	for _, obj := range redataslice {
		objMap, ok := obj.(map[string]interface{})

		if !ok {
			logrus.Error("cannot convert RE interface{} to type map[string]interface{}")
		}

		// The below if statements check on the data type to see if an upload is possible
		if res, ok := objMap["id"].(float64); ok {
			CurId = res
			IdProceed = true
		}

		if res, ok := objMap["description"].(string); ok {
			CurDesc = res
			DescProceed = true
		}

		if res, ok := objMap["price"].(float64); ok {
			CurPrice = res
			PriceProceed = true
		}

		if res, ok := objMap["lien"].(float64); ok {
			CurLien = res
			LienProceed = true
		}

		if res, ok := objMap["rate"].(float64); ok {
			CurRate = res
			RateProceed = true
		}

		if res, ok := objMap["years"].(float64); ok {
			CurYears = res
			YearsProceed = true
		}

		if res, ok := objMap["value"].(float64); ok {
			CurValue = res
			ValueProceed = true
		}

		if res, ok := objMap["monthsleft"].(float64); ok {
			CurMonthsLeft = res
			MonthsLeftProceed = true
		}

		if res, ok := objMap["type"].(string); ok {
			CurType = res
			TypeProceed = true
		}

		if res, ok := objMap["ownership"].(float64); ok {
			CurOwnership = res
			OwnershipProceed = true
		}

		if res, ok := objMap["date"].(string); ok {
			CurDate = res
			DateProceed = true
		}

		tenantsdataslice, ok := objMap["tenants"].([]interface{})

		if !ok {
			logrus.Error("Cannot convert tenant JSON objects")
			logrus.Warn("It is possible this property does not have any tenants")
		}

		var CurName string
		NameProceed := false
		var CurStart string
		StartProceed := false
		var CurEnd string
		EndProceed := false
		var CurUnit string
		UnitProceed := false
		var CurCurTenant bool
		CurTenantProceed := false
		var CurTenantId float64
		TenantIdProceed := false
		var CurSubId float64
		SubIdProceed := false

		var TempTenants []Tenant

		// Tenants data loop
		for _, obj2 := range tenantsdataslice {
			objMap2, ok := obj2.(map[string]interface{})

			if !ok {
				logrus.Error("cannot convert tenants interface{} to type map[string] interface{}")
			}

			// The below if statements check on the data type
			if res, ok := objMap2["name"].(string); ok {
				CurName = res
				NameProceed = true
			}

			if res, ok := objMap2["leasestart"].(string); ok {
				CurStart = res
				StartProceed = true
			}

			if res, ok := objMap2["leaseend"].(string); ok {
				CurEnd = res
				EndProceed = true
			}

			if res, ok := objMap2["unit"].(string); ok {
				CurUnit = res
				UnitProceed = true
			}

			if res, ok := objMap2["currenttenant"].(bool); ok {
				CurCurTenant = res
				CurTenantProceed = true
			}

			if res, ok := objMap2["id"].(float64); ok {
				CurTenantId = res
				TenantIdProceed = true
			}

			if res, ok := objMap2["subid"].(float64); ok {
				CurSubId = res
				SubIdProceed = true
			}

			expensesdataslice, ok := objMap2["expenses"].([]interface{})

			if !ok {
				logrus.Error("Cannot convert expenses JSON objects")
				logrus.Warn("It is possible this tenant does not have any expenses")
			}

			var CurExpenseDescription string
			ExpDescProceed := false
			var CurExpenseAmount float64
			ExpAmtProceed := false
			var CurExpenseDate string
			ExpDateProceed := false

			var TempExpenses []Expense

			for _, obj3 := range expensesdataslice {
				objMap3, ok := obj3.(map[string]interface{})

				if !ok {
					logrus.Error("cannot convert tenant expenses interface{} to map[string]interface{}")
				}

				if res, ok := objMap3["description"].(string); ok {
					CurExpenseDescription = res
					ExpDescProceed = true
				}

				if res, ok := objMap3["amount"].(float64); ok {
					CurExpenseAmount = res
					ExpAmtProceed = true
				}

				if res, ok := objMap3["date"].(string); ok {
					CurExpenseDate = res
					ExpDateProceed = true
				}

				if ExpDescProceed && ExpAmtProceed && ExpDateProceed {
					TempExpenses = append(TempExpenses, Expense{Description: CurExpenseDescription, Amount: CurExpenseAmount, Date: CurExpenseDate})
				}
			}

			revenuesdataslice, ok := objMap2["revenues"].([]interface{})

			if !ok {
				logrus.Error("Cannot convert revenues JSON objects")
				logrus.Warn("It is possible this tenant does not have any revenues")
			}

			var CurRevenueDescription string
			RevDescProceed := false
			var CurRevenueAmount float64
			RevAmtProceed := false
			var CurRevenueDate string
			RevDateProceed := false

			var TempRevenues []Revenue

			for _, obj4 := range revenuesdataslice {
				objMap4, ok := obj4.(map[string]interface{})

				if !ok {
					logrus.Error("cannot convert tenant revenues to interface{}")
				}

				if res, ok := objMap4["description"].(string); ok {
					CurRevenueDescription = res
					RevDescProceed = true
				}

				if res, ok := objMap4["amount"].(float64); ok {
					CurRevenueAmount = res
					RevAmtProceed = true
				}

				if res, ok := objMap4["date"].(string); ok {
					CurRevenueDate = res
					RevDateProceed = true
				}

				if RevDescProceed && RevAmtProceed && RevDateProceed {
					TempRevenues = append(TempRevenues, Revenue{Description: CurRevenueDescription, Amount: CurRevenueAmount, Date: CurRevenueDate})
				}
			}

			if NameProceed && StartProceed && EndProceed && UnitProceed && CurTenantProceed && TenantIdProceed && SubIdProceed {
				TempTenants = append(TempTenants, Tenant{Name: CurName, LeaseStart: CurStart, LeaseEnd: CurEnd, Unit: CurUnit, CurrentTenant: CurCurTenant, Id: CurTenantId, SubId: CurSubId, Expenses: TempExpenses, Revenues: TempRevenues})
			}
		}

		// Building Expenses Loop
		bldgexpdataslice, ok := objMap["buildingexpenses"].([]interface{})

		if !ok {
			logrus.Error("Cannot convert building expenses JSON objects")
			logrus.Warn("It is possible this building does not have any expenses")
		}

		var CurBldgExpDesc string
		BldgExpDescProceed := false
		var CurBldgExpAmt float64
		BldgExpAmtProceed := false
		var CurBldgExpDate string
		BldgExpDateProceed := false

		var TempBuildingExpenses []Expense

		for _, obj5 := range bldgexpdataslice {
			objMap5, ok := obj5.(map[string]interface{})

			if !ok {
				logrus.Error("cannot convert building expenses to interface{}")
			}

			if res, ok := objMap5["description"].(string); ok {
				CurBldgExpDesc = res
				BldgExpDescProceed = true
			}

			if res, ok := objMap5["amount"].(float64); ok {
				CurBldgExpAmt = res
				BldgExpAmtProceed = true
			}

			if res, ok := objMap5["date"].(string); ok {
				CurBldgExpDate = res
				BldgExpDateProceed = true
			}

			if BldgExpDescProceed && BldgExpAmtProceed && BldgExpDateProceed {
				TempBuildingExpenses = append(TempBuildingExpenses, Expense{Description: CurBldgExpDesc, Amount: CurBldgExpAmt, Date: CurBldgExpDate})
			}
		}

		// Building Revenues Loop
		bldgrevsdataslice, ok := objMap["buildingrevenues"].([]interface{})

		if !ok {
			logrus.Error("Cannot convert building revenues JSON Objects")
			logrus.Warn("It is possible this building does not have any revenues")
		}

		var CurBldgRevDesc string
		BldgRevDescProceed := false
		var CurBldgRevAmt float64
		BldgRevAmtProceed := false
		var CurBldgRevDate string
		BldgRevDateProceed := false

		var TempBuildingRevenues []Revenue

		for _, obj6 := range bldgrevsdataslice {
			objMap6, ok := obj6.(map[string]interface{})

			if !ok {
				logrus.Error("cannot convert building revenues to interface{}")
			}

			if res, ok := objMap6["description"].(string); ok {
				CurBldgRevDesc = res
				BldgRevDescProceed = true
			}

			if res, ok := objMap6["amount"].(float64); ok {
				CurBldgRevAmt = res
				BldgRevAmtProceed = true
			}

			if res, ok := objMap6["date"].(string); ok {
				CurBldgRevDate = res
				BldgRevDateProceed = true
			}

			if BldgRevDescProceed && BldgRevAmtProceed && BldgRevDateProceed {
				TempBuildingRevenues = append(TempBuildingRevenues, Revenue{Description: CurBldgRevDesc, Amount: CurBldgRevAmt, Date: CurBldgRevDate})
			}

		}

		if IdProceed && DescProceed && PriceProceed && LienProceed && RateProceed && YearsProceed && ValueProceed && MonthsLeftProceed && TypeProceed && OwnershipProceed && DateProceed {
			CurrentProperties = append(CurrentProperties, RealEstatePost{Id: CurId, Description: CurDesc, Price: CurPrice, Lien: CurLien, Rate: CurRate, Years: CurYears, Value: CurValue, MonthsLeft: CurMonthsLeft, Type: CurType, Ownership: CurOwnership, Date: CurDate, Tenants: TempTenants, BuildingExpenses: TempBuildingExpenses, BuildingRevenues: TempBuildingRevenues})
		}

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

	TempProperties = nil

	if CurrentProperties == nil {
		TempProperties = append(TempProperties, RealEstatePost{Id: 0, Description: propertyData.Description, Price: propertyData.Price, Lien: propertyData.Lien, Rate: propertyData.Rate, Years: propertyData.Years, Value: propertyData.Value, MonthsLeft: propertyData.MonthsLeft, Type: propertyData.Type, Ownership: propertyData.Ownership, Date: propertyData.Date})
		CurrentProperties = TempProperties

		file, _ := json.MarshalIndent(CurrentProperties, "", " ")

		_ = ioutil.WriteFile("Properties.json", file, 0644)
	} else {

		NewId := float64(len(CurrentProperties))

		IdAvailable := false
		IdInBatch := false

		for !IdAvailable {
			IdInBatch = false
			for _, property := range CurrentProperties {
				if NewId == property.Id {
					IdInBatch = true
				}
			}

			if !IdInBatch {
				IdAvailable = true
			} else {
				NewId++
			}
		}

		CurrentProperties = append(CurrentProperties, RealEstatePost{Id: NewId, Description: propertyData.Description, Price: propertyData.Price, Lien: propertyData.Lien, Rate: propertyData.Rate, Years: propertyData.Years, Value: propertyData.Value, MonthsLeft: propertyData.MonthsLeft, Type: propertyData.Type, Ownership: propertyData.Ownership, Date: propertyData.Date})

		file, _ := json.MarshalIndent(CurrentProperties, "", " ")

		_ = ioutil.WriteFile("Properties.json", file, 0644)
	}

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

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

			// Assigning these variables early because they are not sent in the front end
			tenants := CurrentProperties[property].Tenants
			buildingexpenses := CurrentProperties[property].BuildingExpenses

			logrus.Info(tenants)

			CurrentProperties[property].Id = propertyData.Id
			CurrentProperties[property].Description = propertyData.Description
			CurrentProperties[property].Price = propertyData.Price
			CurrentProperties[property].Lien = propertyData.Lien
			CurrentProperties[property].Rate = propertyData.Rate
			CurrentProperties[property].Years = propertyData.Years
			CurrentProperties[property].Value = propertyData.Value
			CurrentProperties[property].MonthsLeft = propertyData.MonthsLeft
			CurrentProperties[property].Type = propertyData.Type
			CurrentProperties[property].Ownership = propertyData.Ownership
			CurrentProperties[property].Date = propertyData.Date
			CurrentProperties[property].Tenants = tenants
			CurrentProperties[property].BuildingExpenses = buildingexpenses

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
		CurrentProperties[0].Ownership = 0
		CurrentProperties[0].Date = "10/01/2022"

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

func addTenant(w http.ResponseWriter, request *http.Request) {
	// Enabling Cors
	enableCors(&w)

	// Sets the appropriate data type for JSON data transfer
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Returns the status of the local host
	w.WriteHeader(http.StatusOK)

	// Initializing API variables
	var propertyData RealEstatePost

	// Initializes the JSON decoder
	decoder := json.NewDecoder(request.Body)

	// Uses the decoder to decode the response to tenant data at its memory location
	decoder.Decode(&propertyData)

	logrus.Info(propertyData)

	for property := range CurrentProperties {
		if propertyData.Id == CurrentProperties[property].Id {
			logrus.Info(propertyData)
			CurrentProperties[property].Tenants = append(CurrentProperties[property].Tenants, propertyData.Tenants[0])
		}
	}

	file, _ := json.MarshalIndent(CurrentProperties, "", " ")

	_ = ioutil.WriteFile("Properties.json", file, 0644)
}
