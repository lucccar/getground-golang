package model

import (
	"encoding/json"
	"io"
	"log"
)

type Response struct {
	Status  int
	Message string
	Data    string
	Name                string `json:"name"`
	ID        int64 `json:"id"`
	Capacity  int   `json:"capacity"`
	Guests string `json:"guests"`
	SeatsEmpty int `json:"seats_empty"`
}

type Body struct {
	Name                string `json:"name" db:"name"`
	Table               int    `json:"table" db:"table"`
	Accompanying_guests int    `json:"accompanying_guests" db:"accompanyingGuests"`
	ID        int64 `json:"id" db:"id"`
	Capacity  int   `json:"capacity" db:"capacity"`
	FreeSeats int   `json:"free_seats" db:"freeseats"`
	
}

func (body *Body) GetRequestBody(rBody io.Reader) error {
	err := json.NewDecoder(rBody).Decode(body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}