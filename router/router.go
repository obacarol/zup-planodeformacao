package router

import (
	"planodeformacao-upgrade/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/account/{id}", middleware.GetAccount).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/accounts", middleware.GetAllAccounts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newaccount", middleware.CreateAccount).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/account/{id}", middleware.UpdateAccount).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteaccount/{id}", middleware.DeleteAccount).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/transaction", middleware.DoTransaction).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/statement/{id}", middleware.GetRecordsByID).Methods("GET", "OPTIONS")

	return router
}
