// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	_ "github.com/go-sql-driver/mysql"
// )

// type table struct {
// 	ID       string
// 	Capacity int
// }

// func addTable(context *gin.Context) {
// 	var newTable table
// 	if err := context.BindJSON(&newTable); err != nil {
// 		return
// 	}
// }

// func main() {
// 	// init mysql.
// 	db, err := sql.Open("mysql", "user:password@/getground")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// router
// 	router := mux.NewRouter()
// 	router.HandleFunc("/getEmployee", controller.AllEmployee).Methods("GET")
// 	router.HandleFunc("/insertEmployee", controller.InsertEmployee).Methods("POST")
// 	http.Handle("/", router)
// 	fmt.Println("Connected to port 1234")
// 	log.Fatal(http.ListenAndServe(":3000", router))

// 	// ping
// 	// http.HandleFunc("/ping", handlerPing)
// 	// http.HandleFunc("/tables", handlerPing)
// 	// http.ListenAndServe(":3000", nil)
// }

// func handlerPing(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "pong\n")
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	controller "github.com/getground/tech-tasks/backend/cmd/Controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tables", controller.AddTable).Methods("POST")
	router.HandleFunc("/guest_list/{name}", controller.AddGuests).Methods("POST")
	router.HandleFunc("/guest_list", controller.GetGuestList).Methods("GET")
	router.HandleFunc("/guests/{name}", controller.GuestArrives).Methods("PUT")
	router.HandleFunc("/guests/{name}", controller.GuestLeaves).Methods("DELETE")
	router.HandleFunc("/guests", controller.GetArrivedGuests).Methods("GET")
	router.HandleFunc("/seats_empty", controller.GetSeatsEmpty).Methods("GET")


	http.Handle("/", router)
	fmt.Println("Connected to port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
