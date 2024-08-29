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
