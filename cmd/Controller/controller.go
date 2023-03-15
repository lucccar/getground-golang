package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	config "github.com/getground/tech-tasks/backend/cmd/Config"
	model "github.com/getground/tech-tasks/backend/cmd/Model"
	utils "github.com/getground/tech-tasks/backend/cmd/Utils"
)

func AddTable(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	capacityString := r.FormValue("capacity")
	capacity, err := strconv.Atoi(capacityString)
	if err != nil {
		return
	}

	id, _ := utils.GenerateID()

	newTable := &model.Table{
		Capacity:  capacity,
		ID:        id,
		FreeSeats: capacity,
	}
	_, err = db.Exec("INSERT INTO tables(ID, capacity, freeseats) VALUES(?)", newTable)

	if err != nil {
		log.Print(err)
		return
	}

	tableJson, _ := json.Marshal(newTable)
	response.Data = string(tableJson)
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func AddGuests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guestName := vars["name"]

	var response model.Response
	var existingTable model.Table

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	tableString := r.FormValue("table")
	accompanyingGuestsString := r.FormValue("accompanying_guests")

	tableID, err := strconv.Atoi(tableString)
	if err != nil {
		return
	}
	accompanyingGuests, err2 := strconv.Atoi(accompanyingGuestsString)
	if err2 != nil {
		return
	}

	tableRow := db.QueryRow("SELECT FROM tables(ID) VALUES(?)", tableID)

	err = tableRow.Scan(&existingTable.ID, &existingTable.Capacity, &existingTable.FreeSeats)
	if err != nil {
		return
	}
	if existingTable.Capacity < accompanyingGuests {
		response.Status = 400
		response.Message = "Insuficient space at the table"
	}

	freeSeats := existingTable.Capacity - accompanyingGuests
	_, err = db.Exec("UPDATE tables SET freeseats = (?) WHERE id = (?)", freeSeats, tableID)
	if err != nil {
		response.Status = 500
		response.Message = "Something wrong happened, try again!"
	}

	if response.Status == 0 {
		response.Data = guestName
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetGuests(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var existingTable model.Table

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	tableString := r.FormValue("table")
	accompanyingGuestsString := r.FormValue("accompanying_guests")

	tableID, err := strconv.Atoi(tableString)
	if err != nil {
		return
	}
	accompanyingGuests, err2 := strconv.Atoi(accompanyingGuestsString)
	if err2 != nil {
		return
	}

	tableRow := db.QueryRow("SELECT FROM tables(ID) VALUES(?)", tableID)

	err = tableRow.Scan(&existingTable.ID, &existingTable.Capacity, &existingTable.FreeSeats)
	if err != nil {
		return
	}
	if existingTable.Capacity < accompanyingGuests {
		response.Status = 400
		response.Message = "Insuficient space at the table"
	}

	freeSeats := existingTable.Capacity - accompanyingGuests
	_, err = db.Exec("UPDATE tables SET freeseats = (?) WHERE id = (?)", freeSeats, tableID)
	if err != nil {
		response.Status = 500
		response.Message = "Something wrong happened, try again!"
	}

	if response.Status == 0 {
		response.Data = guestName
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
