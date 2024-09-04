package client

import (
	"io"
	"net/http"
)

func GetCardData() ([]byte, error) {
	resp, err := http.Get("https://api.scryfall.com/cards/named?fuzzy=lightning%20bolt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
