package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thiagosena/go-finance/adapter/http/actuator"
	"github.com/thiagosena/go-finance/adapter/http/transaction"
	"net/http"
)

// Init function to start all routes of the application
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":7000", nil)
}
