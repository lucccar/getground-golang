package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	config "github.com/getground/tech-tasks/backend/cmd/Config"
	model "github.com/getground/tech-tasks/backend/cmd/Model"
)

func AddTable(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var body model.Body

	log.SetOutput(os.Stdout)

	err := body.GetRequestBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		
		response.Status = 401
		response.Message = "Something went wrong!"
		json.NewEncoder(w).Encode(response)
		return 
	}

	newTable := &model.Table{
		Capacity:  body.Capacity,
		FreeSeats: body.Capacity,
	}
	newTable.AddTable()

	tableJson, _ := json.Marshal(newTable)

	response.Data = string(tableJson)
	response.ID = newTable.ID
	response.Capacity = newTable.Capacity
	response.Status = 200
	response.Message = "Insert data successfully"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	json.NewEncoder(w).Encode(response)
}

func AddGuests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guestName := vars["name"]

	var response model.Response
	var existingTable model.Table
	var body model.Body

	err := body.GetRequestBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		
		response.Status = 401
		response.Message = "Something went wrong!"
		json.NewEncoder(w).Encode(response)
		return 
	}

	newGuest := &model.GuestExpected{
		Name: guestName,
		Table: body.Table,
		Accompanying_guests: body.Accompanying_guests,
	}
	
	err = existingTable.GetTableByID(body.Table)
	if err != nil {
		response.Message = err.Error()
	}

	if existingTable.Capacity < body.Accompanying_guests {
		response.Status = 400
		response.Message = "Insuficient space at the table"
	} else {
		err = newGuest.AddExpectedGuest()
		if err != nil {
			response.Status = 500
			response.Message = err.Error()
		}
	}

	freeSeats := existingTable.Capacity - body.Accompanying_guests
	err = existingTable.UpdateFreeSeats(freeSeats)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
	}

	if response.Status == 0 {
		response.Name = guestName
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetGuestList(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var expectedGuest model.GuestExpected
	var expectedGuests []model.GuestExpected

	expectedGuests, err := expectedGuest.GetAllGuestList()
	if err != nil {
		response.Message = err.Error()
	}

	if response.Status == 0 {
		guestsJSON, _ := json.Marshal(expectedGuests)
		response.Guests = string(guestsJSON)
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GuestArrives(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guestName := vars["name"]
	
	var body model.Body
	var response model.Response
	var existingGuest model.GuestExpected

	var existingTable model.Table

	err := body.GetRequestBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		response.Status = 401
		response.Message = "Something went wrong!"
		json.NewEncoder(w).Encode(response)
		return 
	}

	err = existingGuest.GetExpectedGuestByName(guestName)
	if err != nil {
		response.Message = err.Error()
	}

	if existingGuest.Accompanying_guests > body.Accompanying_guests {
		arrivedGuest, _ := existingGuest.AddArrivedGuest()
		response.Name = arrivedGuest.Name
		response.Status = 200
	} else {
		err = existingTable.GetTableByID(existingGuest.Table)
		fmt.Println("existingTable.Capacity ==> ", existingTable.Capacity)
		if err != nil {
			response.Status = 500
			response.Message = "Something wrong happened, try again!"
		}
		if existingTable.FreeSeats > body.Accompanying_guests {
			newFreeSeats := existingTable.FreeSeats - body.Accompanying_guests
			go existingTable.UpdateFreeSeats(newFreeSeats)
			arrivedGuest, _ := existingGuest.AddArrivedGuest()
			response.Name = arrivedGuest.Name
			response.Status = 200
		} else {
			response.Status = 400
			response.Message = "Too many accompaning guests!"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GuestLeaves(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guestName := vars["name"]

	var response model.Response
	var table model.Table
	var guest model.Guest

	err := guest.GetExpectedGuestByName(guestName)
	if err != nil {
		response.Status = 500
		response.Message = "Something wrong happened, try again!"
	}

	go guest.DeleteGuestByName()

	err = table.GetTableByID(guest.Table)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
	} else {
		newFreeSeats := table.FreeSeats + guest.Accompanying_guests 
		go table.UpdateFreeSeats(newFreeSeats)
	}

	if response.Status == 0 {
		response.Status = 204
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetArrivedGuests(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	var arrivedGuest model.Guest

	db := config.Connect()
	defer db.Close()

	arrivedGuests, err := arrivedGuest.GetAllGuests()
	if err != nil {
		response.Status = 400
		response.Message = err.Error()
	}

	if response.Status == 0 {
		arrivedGuestsJSON, _ := json.Marshal(arrivedGuests)
		response.Guests = string(arrivedGuestsJSON)
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func GetSeatsEmpty(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	freeSeats, err := model.GetFreeSeats()
	if err != nil {
		response.Message = err.Error()
	}

	if response.Status == 0 {
		seatsEmptyJSON, _ := json.Marshal(freeSeats)
		response.SeatsEmpty, _ = strconv.Atoi(string(seatsEmptyJSON))
		response.Status = 200
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}