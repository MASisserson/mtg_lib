package model

import (
	"encoding/json"
	"masisserson/mtg-lib/utils"
)

type ColorSet struct {
	Colors []string
	White  bool
	Blue   bool
	Black  bool
	Red    bool
	Green  bool
}

func (c *ColorSet) UnmarshalJSON(data []byte) error {
	var colorSlice []string
	if err := json.Unmarshal(data, &colorSlice); err != nil {
		return err
	}

	colorsMap := utils.SliceToMap(colorSlice)
	c.Colors = colorSlice
	c.White = colorsMap["W"]
	c.Blue = colorsMap["U"]
	c.Black = colorsMap["B"]
	c.Red = colorsMap["R"]
	c.Green = colorsMap["G"]

	return nil
}

func (c *ColorSet) SetColors() {
	if c.White {
		c.Colors = append(c.Colors, "W")
	}
	if c.Blue {
		c.Colors = append(c.Colors, "U")
	}
	if c.Black {
		c.Colors = append(c.Colors, "B")
	}
	if c.Red {
		c.Colors = append(c.Colors, "R")
	}
	if c.Green {
		c.Colors = append(c.Colors, "G")
	}
}

func (c *ColorSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Colors []string `json:"colors"`
		White  bool     `json:"white"`
		Blue   bool     `json:"blue"`
		Black  bool     `json:"black"`
		Red    bool     `json:"red"`
		Green  bool     `json:"green"`
	}{
		Colors: c.Colors,
		White:  c.White,
		Blue:   c.Blue,
		Black:  c.Black,
		Red:    c.Red,
		Green:  c.Green,
	})
}

func (c *Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        *int   `json:"id"`
		Name      string `json:"name"`
		WizID     string `json:"wiz_id"`
		Set       string `json:"set"`
		ImageURIs struct {
			Small      string `json:"small"`
			Normal     string `json:"normal"`
			Large      string `json:"large"`
			PNG        string `json:"png"`
			ArtCrop    string `json:"art_crop"`
			BorderCrop string `json:"border_crop"`
		} `json:"image_uris"`
		OracleText string   `json:"oracle_text"`
		Type       string   `json:"type_line"`
		Colors     ColorSet `json:"color_set"`
		ManaValue  float32  `json:"mana_value"`
		ManaCost   string   `json:"mana_cost"`
		Power      string   `json:"power"`
		Toughness  string   `json:"toughness"`
		Loyalty    string   `json:"loyalty"`
		Quantity   *int     `json:"quantity"`
	}{
		ID:         c.ID,
		Name:       c.Name,
		WizID:      c.WizID,
		Set:        c.Set,
		ImageURIs:  c.ImageURIs,
		OracleText: c.OracleText,
		Type:       c.Type,
		Colors:     c.Colors,
		ManaValue:  c.ManaValue,
		ManaCost:   c.ManaCost,
		Power:      c.Power,
		Toughness:  c.Toughness,
		Loyalty:    c.Loyalty,
		Quantity:   c.Quantity,
	})
}

// Consider changing ID and Quantity json settings
type Card struct {
	ID        *int
	Name      string `json:"name"`
	WizID     string `json:"id"`
	Set       string `json:"set"`
	ImageURIs struct {
		Small      string `json:"small"`
		Normal     string `json:"normal"`
		Large      string `json:"large"`
		PNG        string `json:"png"`
		ArtCrop    string `json:"art_crop"`
		BorderCrop string `json:"border_crop"`
	} `json:"image_uris"`
	OracleText string   `json:"oracle_text"`
	Type       string   `json:"type_line"`
	Colors     ColorSet `json:"colors"`
	ManaValue  float32  `json:"cmc"`
	ManaCost   string   `json:"mana_cost"`
	Power      string   `json:"power"`
	Toughness  string   `json:"toughness"`
	Loyalty    string   `json:"loyalty"`
	Quantity   *int
}
