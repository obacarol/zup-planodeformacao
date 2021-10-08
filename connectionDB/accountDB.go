package connectionDB

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"planodeformacao-upgrade/models"
)

func GetAllAccounts() ([]models.Account, error) {
	db := DbConn()
	defer db.Close()
	var accounts []models.Account
	sqlStatement := `SELECT * FROM account`
	res, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
	}
	defer res.Close()
	for res.Next() {
		var account models.Account
		err = res.Scan(&account.ID, &account.Name, &account.Cpf, &account.Creation_date, &account.Balance_account)
		if err != nil {
			log.Printf("Error to scan the row. %v", err)
		}
		accounts = append(accounts, account)
	}
	return accounts, err
}

func GetAccount(id int64) (models.Account, error) {
	db := DbConn()
	defer db.Close()
	var account models.Account
	sqlStatement := `SELECT * FROM account WHERE id=$1`
	res := db.QueryRow(sqlStatement, id)
	err := res.Scan(&account.ID, &account.Name, &account.Cpf, &account.Creation_date, &account.Balance_account)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return account, err
	case nil:
		return account, nil
	default:
		log.Printf("Error to scan the row. %v", err)
	}
	return account, err
}

func InsertAccount(account models.Account) int64 {
	db := DbConn()
	defer db.Close()
	sqlStatement := `INSERT INTO account (name, cpf) VALUES ($1, $2) RETURNING id`
	var id int64
	err := db.QueryRow(sqlStatement, account.Name, account.Cpf).Scan(&id)
	if err != nil {
		log.Printf("Error to execute the query. %v", err)
	}
	fmt.Printf("Id %v inserted successfully.\n", id)
	return id
}

func UpdateAccount(id int64, account models.Account) error {
	db := DbConn()
	defer db.Close()
	sqlStatement := `UPDATE account SET name=$2, cpf=$3 WHERE id=$1`
	_, err := db.Exec(sqlStatement, id, account.Name, account.Cpf)
	if err != nil {
		log.Printf("Error to execute the query. %v", err)
		return err
	}
	fmt.Printf("Id %v updated successfully.\n", id)
	return nil
}

func DeleteAccount(id int64) error {
	db := DbConn()
	defer db.Close()
	sqlStatement := `DELETE FROM account WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Printf("Error to execute the query. %v", err)
		return err
	}
	fmt.Printf("Id %v deleted successfully.\n", id)
	return nil
}

func DoTransaction(record models.Records) error {
	db := DbConn()
	defer db.Close()
	err := processAccount(record)
	if err != nil {
		log.Printf("Error to process the account. %v", err)
		return err
	} else {
		sqlStatement := `INSERT INTO records (id_account_from, id_account_to, transaction_value, transaction_type) VALUES ($1, $2, $3, $4)`
		_, err := db.Exec(sqlStatement, record.ID_account_from, record.ID_account_to, record.Transaction_value, record.Transaction_type)
		if err != nil {
			log.Printf("Error to execute the query. %v", err)
			return err
		}
		fmt.Printf("%v executed successfully.\n", record.Transaction_type)
		return err
	}
}

func processAccount(record models.Records) error {
	switch record.Transaction_type {
	case "Pay in":
		err := updateBalance(record.ID_account_from, record.Transaction_value, true)
		return err
	case "Withdrawal":
		err := checkBalance(record.ID_account_from, record.Transaction_value)
		if err != nil {
			return err
		} else {
			err := updateBalance(record.ID_account_from, record.Transaction_value, false)
			return err
		}
	case "Transfer":
		//Checa se tem saldo suficiente na conta origem
		err := checkBalance(record.ID_account_from, record.Transaction_value)
		if err != nil {
			return err
		}
		//idAccountTo recebe o valor de ID_account_to ou zero em caso de nulo
		idAccountTo := record.ID_account_to.ValueOrZero()
		//Verifica se a conta destino existe
		_, err = GetAccount(idAccountTo)
		if err != nil {
			return err
		}
		//Caso a conta origem possua fundos, e a conta destino seja válida, as contas terão seus saldos atualizados
		err = updateBalance(idAccountTo, record.Transaction_value, true)
		if err != nil {
			return err
		}
		err = updateBalance(record.ID_account_from, record.Transaction_value, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateBalance(idClient int64, value float64, credit bool) error {
	db := DbConn()
	defer db.Close()
	if !credit {
		value = (value * -1)
	}
	sqlStatement := `UPDATE account SET balance_account=($2 + balance_account) WHERE id=$1`
	_, err := db.Exec(sqlStatement, idClient, value)
	if err != nil {
		log.Printf("Error to execute the query. %v", err)
	}
	return err
}

func checkBalance(idClient int64, value float64) error {
	client, err := GetAccount(idClient)
	if err != nil {
		log.Printf("Error to get an account. %v", err)
		return err
	}
	if client.Balance_account-value < 0 {
		return errors.New("Insufficient funds")
	}
	return nil
}

func Statement(id int64) ([]models.Records, error) {
	db := DbConn()
	defer db.Close()
	var records []models.Records
	sqlStatement := `SELECT r.id_account_from, r.id_account_to, r.transaction_date, r.transaction_type, r.transaction_value FROM records AS r WHERE r.id_account_from=$1 OR r.id_account_to=$1`
	res, err := db.Query(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	for res.Next() {
		var record models.Records
		err = res.Scan(&record.ID_account_from, &record.ID_account_to, &record.Transaction_date, &record.Transaction_type, &record.Transaction_value)
		if err != nil {
			log.Printf("Error to scan the row. %v", err)
		}
		records = append(records, record)
	}
	return records, err
}
