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

	"github.com/Nhanderu/brdoc"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type responseAccount struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := connectionDB.GetAllAccounts()
	if err != nil {
		log.Printf("Error to get all accounts. %v", err)
		http.Error(w, "Error to get all accounts", 500)
	}
	json.NewEncoder(w).Encode(accounts)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error to convert ID to integer. %v", err)
		http.Error(w, "ID must be valid", 500)
	} else {
		account, err := connectionDB.GetAccount(int64(id))
		if err != nil {
			log.Printf("Error to get an account. %v", err)
			http.Error(w, "ID must be valid", 500)
		} else {
			json.NewEncoder(w).Encode(account)
		}
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var err, account = processAccount(r.Body)
	if err != "" {
		log.Println(err)
		http.Error(w, err, 500)
	} else {
		insertID := connectionDB.InsertAccount(account)
		res := responseAccount{
			ID:      insertID,
			Message: "Account created successfully.",
		}
		json.NewEncoder(w).Encode(res)
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error to convert ID to integer. %v", err)
		http.Error(w, "ID must be valid", 500)
	} else {
		err, account := processAccount(r.Body)
		if err != "" {
			log.Println(err)
			http.Error(w, err, 500)
		} else {
			var err = connectionDB.UpdateAccount(int64(id), account)
			if err != nil {
				log.Printf("Error to update the account. %v", err)
				http.Error(w, "Error to update the account", 500)
			} else {
				msg := fmt.Sprintf("Account %v updated successfully.", id)
				res := responseAccount{
					ID:      int64(id),
					Message: msg,
				}
				json.NewEncoder(w).Encode(res)
			}
		}
	}
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Error to convert ID to integer. %v", err)
		http.Error(w, "ID must be valid", 500)
	} else if !IsValidID(int64(id)) {
		log.Printf("ID not found. %v", err)
		http.Error(w, "ID not found", 500)
	} else {
		err = connectionDB.DeleteAccount(int64(id))
		if err != nil {
			log.Printf("Error to delete the account. %v", err)
			http.Error(w, "Error to delete the account.", 500)
		} else {
			msg := fmt.Sprintf("Account %v deleted successfully.", id)
			res := responseAccount{
				ID:      int64(id),
				Message: msg,
			}
			json.NewEncoder(w).Encode(res)
		}
	}
}

func processAccount(body io.ReadCloser) (string, models.Account) {
	var account models.Account
	err := json.NewDecoder(body).Decode(&account)
	cpf := strconv.Itoa(account.Cpf)
	msgError := ""
	if err != nil {
		msgError = "Invalid JSON"
	} else if account.Name == "" {
		msgError = "Name needs to be informed."
	} else if !brdoc.IsCPF(cpf) {
		msgError = "CPF must be valid."
	} else if account.Creation_date != "" {
		msgError = "'Transaction Date' needs to be null."
	} else if account.Balance_account != 0 {
		msgError = "'Balance account' initial must be zero."
	} else if account.ID != 0 {
		msgError = "ID account is automatically set."
	}
	return msgError, account
}
