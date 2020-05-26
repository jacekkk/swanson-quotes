package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"swanson/models"

	"github.com/gorilla/mux"
)

type PageData struct {
	PageTitle string
	Quotes    []*models.Quote
}

type QuotesResponse struct {
	Count  int             `json:"count" binding:"required"`
	Quotes []*models.Quote `json:"quotes" binding:"required"`
}

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", renderPage)
	router.HandleFunc("/api/quotes", getQuotes).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/quotes/{id}", getQuote).Methods("GET", "OPTIONS")

	return router
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	quotes, err := models.GetAllQuotes()

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response := QuotesResponse{
		Count:  len(quotes),
		Quotes: quotes,
	}

	json.NewEncoder(w).Encode(response)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, _ := mux.Vars(r)["id"]

	quote, err := models.GetQuote(id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(quote)
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template.html"))
	quotes, err := models.GetAllQuotes()

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data := PageData{
		PageTitle: "Ron Swanson says...",
		Quotes:    quotes,
	}
	tmpl.Execute(w, data)
}
