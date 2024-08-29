package handlers

import (
	"encoding/json"
	"html/template"
	"masisserson/mtg-lib/client"
	"masisserson/mtg-lib/controller"
	"masisserson/mtg-lib/model"
	"net/http"
)

// INCOMPLETE - Handler to view a given library. Library specified
// in the path.
func LibViewHandler(w http.ResponseWriter, r *http.Request) {
	body, err := client.GetCardData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var card *model.Card
	if err := json.Unmarshal(body, &card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// save card to database
	if err := controller.CreateCard(*card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// pull individual card from database
	card, err = controller.GetCard("Lightning Bolt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("html/lib_view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
