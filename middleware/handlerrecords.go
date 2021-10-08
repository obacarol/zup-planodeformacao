package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"planodeformacao-upgrade/connectionDB"
	"planodeformacao-upgrade/models"

	"github.com/gorilla/mux"
)

type responseRecord struct {
	Transaction models.Transaction `json:"transaction,omitempty"`
	Message     string             `json:"message,omitempty"`
}

func DoTransaction(w http.ResponseWriter, r *http.Request) {
	var err, record = processRecord(r.Body)
	if err != "" {
		log.Println(err)
		http.Error(w, err, 500)
	} else {
		erro := connectionDB.DoTransaction(record)
		if erro != nil {
			log.Printf("Error to do the transaction. %v", err)
			http.Error(w, "Error to do the transaction", 500)
		} else {
			msg := fmt.Sprintf("%v: $ %v - from ID:%v to ID:%v - executed successfully.", record.Transaction_type, record.Transaction_value, record.ID_account_from, record.ID_account_to)
			res := responseRecord{
				Transaction: record.Transaction_type,
				Message:     msg,
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}

func processRecord(body io.ReadCloser) (string, models.Records) {
	var record models.Records
	err := json.NewDecoder(body).Decode(&record)
	msgError := ""
	if err != nil {
		msgError = "Invalid JSON"
	} else if record.Transaction_date != "" {
		msgError = "'Transaction Date' needs to be null."
	} else if record.Transaction_value <= 0 {
		msgError = "'Transaction value' cannot be less than or equal to zero."
	} else if !isValidTransaction(record.Transaction_type) {
		msgError = "Invalid transaction type"
	} else if !IsValidID(record.ID_account_from) {
		msgError = "ID_account_from is invalid"
	} else if !IsValidID(record.ID_account_to.Int64) && !record.ID_account_to.IsZero() {
		msgError = "ID_account_to is invalid"
	}
	return msgError, record
}

func isValidTransaction(tt models.Transaction) bool {
	switch tt {
	case models.Pay_In, models.Withdrawal, models.Transfer:
		return true
	}
	return false
}

func IsValidID(idClient int64) bool {
	_, err := connectionDB.GetAccount(idClient)
	if err != nil {
		log.Printf("Invalid ID. %v", err)
		return false
	}
	return true
}

func GetRecordsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error to convert ID to integer. %v", err)
		http.Error(w, "ID must be valid", 500)
	} else {
		records, err := connectionDB.Statement(int64(id))
		if err != nil {
			log.Printf("Error to get the statement. %v", err)
			http.Error(w, "ID must be valid", 500)
		} else {
			json.NewEncoder(w).Encode(records)
		}
	}
}
