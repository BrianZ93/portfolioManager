package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	readFile()
	readREFile()
	readDebtFile()
	getPrices()

	// This function must be the last one as it ties up the terminal
	handleRequests()

}

// Function to create the welcome page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, currentTime.Format("01-02-2006"))
	fmt.Fprintf(w, "\nHome Page Initiated")
	fmt.Fprintf(w, "\nhttp://localhost:8081/index to inspect the index page")
	fmt.Fprintf(w, "\nhttp://localhost:8081/portfolio to inspect the portfolio page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/index", indexPage)

	// Equities REST
	http.HandleFunc("/portfolio", retrieveEquities)
	http.HandleFunc("/portfolioadd", addEquity)
	http.HandleFunc("/portfoliomod", modifyEquity)
	http.HandleFunc("/portfoliodel", deleteEquity)

	// Properties REST
	http.HandleFunc("/properties", RERetrieve)
	http.HandleFunc("/propertyadd", addProperty)
	http.HandleFunc("/propertymod", modifyProperty)
	http.HandleFunc("/propertydel", deleteProperty)

	// Debt REST
	http.HandleFunc("/debts", DebtRetrieve)
	http.HandleFunc("/debtadd", addDebt)
	http.HandleFunc("/debtmod", modifyDebt)
	http.HandleFunc("/debtdel", deleteDebt)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Needed to allow our VueJS project access to the backend
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
