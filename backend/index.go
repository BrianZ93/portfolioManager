package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/piquette/finance-go/index"
)

// Initializing global variables
var indexesprice []string
var indexeschange []string
var x string = "x"
var currentTime = time.Now()

// var pulledToday = false

type Index struct {
	Title  string `json:"title"`
	Price  string `json:"price"`
	Change string `json:"change"`
}

type Indexes []Index

func getPrices() {
	// Getting index prices - the first one must be done independently in order to define the "err" variable
	i, err := index.Get("^VIX")
	if err != nil {
		panic(err)
	}

	x = strconv.FormatFloat(i.RegularMarketPrice, 'f', -1, 64)

	indexesprice = append(indexesprice, x)

	x = strconv.FormatFloat(i.RegularMarketChangePercent, 'f', -1, 64)

	indexeschange = append(indexeschange, x)

	symbols := []string{"^GSPC", "^DJI", "^IXIC"}
	iter := index.List(symbols)

	for iter.Next() {
		q := iter.Index()

		x = strconv.FormatFloat(q.RegularMarketPrice, 'f', -1, 64)

		indexesprice = append(indexesprice, x)

		x = strconv.FormatFloat(q.RegularMarketChangePercent, 'f', -1, 64)

		indexeschange = append(indexeschange, x)
	}

	if iter.Err() != nil {
		panic(err)
	}

}

// Function to retrieve market index data
func indexPage(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS
	enableCors(&w)

	// Setting up the JSON objects
	w.Header().Set("Content-Type", "application/json")

	x_indexes := Indexes{
		Index{Title: "VIX", Price: indexesprice[0], Change: indexeschange[0]},
		Index{Title: "GSPC", Price: indexesprice[1], Change: indexeschange[1]},
		Index{Title: "DJI", Price: indexesprice[2], Change: indexeschange[2]},
		Index{Title: "IXIC", Price: indexesprice[3], Change: indexeschange[3]},
	}

	fmt.Println("Index Endpoint Hit")
	json.NewEncoder(w).Encode(x_indexes)
}
