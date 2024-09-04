package controller

import (
	"database/sql"
	"encoding/json"
	"masisserson/mtg-lib/db"
	"masisserson/mtg-lib/model"
	_ "masisserson/mtg-lib/model"
)

func CreateCard(card model.Card) error {
	query := `
		INSERT INTO card
		("name", wiz_id, "set", oracle, "type", white, blue, black, red,
		green, mana_value, mana_cost, power, toughness, loyalty, img)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
		$14, $15, $16)
	`
	jsonb, err := json.Marshal(card.ImageURIs)
	if err != nil {
		return err
	}

	objects := []any{
		card.Name,
		card.WizID,
		card.Set,
		card.OracleText,
		card.Type,
		card.Colors.White,
		card.Colors.Blue,
		card.Colors.Black,
		card.Colors.Red,
		card.Colors.Green,
		card.ManaValue,
		card.ManaCost,
		card.Power,
		card.Toughness,
		card.Loyalty,
		jsonb,
	}

	// need to check if there is already a card copy present

	if err := db.Create(query, objects); err != nil {
		return err
	}

	return nil
}

func extractCardRows(rows *sql.Rows) ([]*model.Card, error) {
	var cards []*model.Card
	for rows.Next() {
		card := &model.Card{}
		imgJSON := []byte{}
		if err := rows.Scan(
			&card.ID,
			&card.Name,
			&card.WizID,
			&card.Set,
			&card.OracleText,
			&card.Type,
			&card.Colors.White,
			&card.Colors.Blue,
			&card.Colors.Black,
			&card.Colors.Red,
			&card.Colors.Green,
			&card.ManaValue,
			&card.ManaCost,
			&card.Power,
			&card.Toughness,
			&card.Loyalty,
			&imgJSON,
			&card.Quantity,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(imgJSON, &card.ImageURIs); err != nil {
			return nil, err
		}
		card.Colors.SetColors()

		cards = append(cards, card)
	}

	return cards, nil
}

func GetCards(name string) ([]*model.Card, error) {
	query := `
		SELECT * FROM card
	`

	rows, err := db.Read(query)
	if err != nil {
		return nil, err
	}

	cards, err := extractCardRows(rows)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
