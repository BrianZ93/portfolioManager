package main

import (
	"fmt"
	"log"
	"net/http"
	// "bytes"
	// "uuid"
	// "flag"
	// "os"
	// "github.com/sirupsen/logrus"
)

func main() {

	readFile()
	getPrices()
	fmt.Println("MM-DD-YYYY: ", currentTime.Format("01-02-2006"))
	handleRequests()

}

// Function to create the welcome page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, currentTime.Format("01-02-2006"))
	fmt.Fprintf(w, "\nHome Page Initiated")
	fmt.Fprintf(w, "\nhttp://localhost:8081/index to inspect the index page")
	fmt.Fprintf(w, "\nhttp://localhost:8081/portfolio to inspect the portfolio page")
}

// Function to create the equities POST page
func equitiesPage(w http.ResponseWriter, r *http.Request) {
	// Enabling CORS
	enableCors(&w)

	fmt.Fprintf(w, "POST API Endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/index", indexPage)
	http.HandleFunc("/equities", equitiesPage)
	http.HandleFunc("/realestate", realEstatePage)
	http.HandleFunc("/portfolio", retrieveEquities)
	http.HandleFunc("/portfolioadd", addEquity)
	http.HandleFunc("/portfoliomod", modifyEquity)
	http.HandleFunc("/portfoliodel", deleteEquity)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Needed to allow our VueJS project access to the backend
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
