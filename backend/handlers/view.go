package handlers

import (
	"encoding/json"
	"masisserson/mtg-lib/controller"
	"net/http"
)

// INCOMPLETE - Handler to view a given library. Library specified
// in the path.
func LibViewHandler(w http.ResponseWriter, r *http.Request) {
	// // pull card data from Scryfall API
	// body, err := client.GetCardData()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // unmarshal the card data into the card struct
	// var card *model.Card
	// if err := json.Unmarshal(body, &card); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // save card to database
	// if err := controller.CreateCard(*card); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// pull individual card from database
	cards, err := controller.GetCards("Lightning Bolt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// might run into trouble because I'm passing pointers rather than
	// actual card data
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cards); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// t, err := template.ParseFiles("html/lib_view.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// if err := t.Execute(w, card); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
