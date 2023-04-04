package model

import (
	"log"

	config "github.com/getground/tech-tasks/backend/cmd/Config"
)

type Table struct {
	ID        int64 `json:"id" db:"id"`
	Capacity  int   `json:"capacity" db:"capacity"`
	FreeSeats int   `json:"free_seats" db:"freeseats"`
}

func (table *Table) AddTable() error {

	db := config.Connect()
	defer db.Close()

	insertadeTable, err := db.Exec("INSERT INTO tables(capacity, freeseats) VALUES(?, ?);", table.Capacity, table.FreeSeats)
	if err != nil {
		log.Print(err)
		return err
	}
	table.ID, _ = insertadeTable.LastInsertId()

	return nil
}

func (table *Table) GetTableByID(tableID int) error {

	db := config.Connect()
	defer db.Close()

	tableRow := db.QueryRow("SELECT id, capacity, freeseats FROM tables WHERE id = ? ;", tableID)
	err := tableRow.Scan(&table.ID, &table.Capacity, &table.FreeSeats)
	if err != nil {
		return err
	}

	return nil
}

func (table Table) UpdateFreeSeats(freeSeats int) error {

	db := config.Connect()
	defer db.Close()

	_, err := db.Exec("UPDATE IGNORE tables SET freeseats = ? WHERE id = ? ;", freeSeats, table.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func GetFreeSeats() (int, error) {
	var seatsEmpty int
	db := config.Connect()
	defer db.Close()

	seats := db.QueryRow("SELECT SUM(freeseats) as seats FROM tables ;")
	err := seats.Scan(&seatsEmpty)
	if err != nil {
		return 0, err
	}

	return seatsEmpty, nil
}