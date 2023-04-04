package model

import (
	"log"
	"time"

	config "github.com/getground/tech-tasks/backend/cmd/Config"
)

type GuestExpected struct {
	ID                  int64  `form:"id" json:"id" db:"id"`
	Name                string `form:"name" json:"name" db:"name"`
	Table               int    `form:"table" json:"table" db:"table"`
	Accompanying_guests int    `form:"accompanying_guests" json:"accompanying_guests" db:"accompanyingGuests"`
}

type Guest struct {
	GuestExpected
	TimeArrived string `form:"time_arrived" json:"time_arrived" db:"timeArrived"`
}

func (expectedGuest *GuestExpected) AddExpectedGuest() error {

	db := config.Connect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO expectedguests(name, `table`, accompanyingGuests) VALUES(?, ?, ?);", 
		expectedGuest.Name, 
		expectedGuest.Table, 
		expectedGuest.Accompanying_guests)
	if err != nil {
		log.Print(err)
		return err
	}


	return nil
}

func (expectedGuest *GuestExpected) AddArrivedGuest() (*Guest, error) {
	db := config.Connect()
	defer db.Close()
	timeNow := time.Now()
	_, err := db.Exec("INSERT INTO guests(name, `table`, accompanyingGuests, timeArrived) VALUES(?, ?, ?, ?);", 
		expectedGuest.Name, 
		expectedGuest.Table, 
		expectedGuest.Accompanying_guests,
		timeNow)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	insertedGuest := &Guest{
		*expectedGuest,
		timeNow.String(),
	}

	return insertedGuest, nil
}

func (expectedGuest *GuestExpected) GetAllGuestList() ([]GuestExpected, error) {
	var expectedGuests []GuestExpected
	db := config.Connect()
	defer db.Close()

	guestsRows, err := db.Query("SELECT name, `table`, accompanyingGuests FROM expectedguests;")
	if err != nil {
		return []GuestExpected{}, err
	}

	for guestsRows.Next() {
		err = guestsRows.Scan(&expectedGuest.Name, &expectedGuest.Table, &expectedGuest.Accompanying_guests )
		if err != nil {
			return []GuestExpected{}, err
		}
		expectedGuests = append(expectedGuests, *expectedGuest)
	}

	return expectedGuests, nil
}

func (guest *Guest) DeleteGuestByName() error {
	db := config.Connect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM guests WHERE name = ? ;", guest.Name)
	if err != nil {
		return err
	}

	return nil
}

func (guest *GuestExpected) GetExpectedGuestByName(name  string) error {
	db := config.Connect()
	defer db.Close()

	guestsRow := db.QueryRow("SELECT name, accompanyingGuests, `table` FROM expectedguests WHERE NAME = ? ", name)
	
	err := guestsRow.Scan(&guest.Name, &guest.Accompanying_guests, &guest.Table)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}

func (guest *Guest) GetAllGuests() ([]Guest, error) {
	var guests []Guest
	db := config.Connect()
	defer db.Close()

	guestsRows, err := db.Query("SELECT name, accompanyingGuests, `table` FROM guests")
	if err != nil {
		return []Guest{}, err
	}

	for guestsRows.Next() {
		err = guestsRows.Scan(&guest.Name, &guest.Accompanying_guests, guest.Table)
		if err != nil {
			return []Guest{}, err
		}
		guests = append(guests, *guest)
	}

	return guests, nil
}
