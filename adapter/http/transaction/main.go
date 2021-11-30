package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/thiagosena/go-finance/model/transaction"
	"github.com/thiagosena/go-finance/util"
	"io/ioutil"
	"net/http"
)

// GetTransactions function to get all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.00,
			Type:      0,
			CreatedAt: util.StringToTime("2021-11-28T22:18:25"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

// CreateTransaction function to create a transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transaction{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Println(res)
}
