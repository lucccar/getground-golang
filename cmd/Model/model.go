package model

type Table struct {
	ID        int64 `form:"id" json:"id" db:"id"`
	Capacity  int   `form:"capacity" json:"capacity" db:"capacity"`
	FreeSeats int   `form:"free_seats" json:"free_seats" db:"freeseats"`
}

type Response struct {
	Status  int
	Message string
	Data    string
}
