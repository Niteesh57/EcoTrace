package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type TransactionRequest struct {
	ShipmentData string `json:"shipmentData"`
}

func submitTransaction(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	json.NewDecoder(r.Body).Decode(&req)

	wallet, err := gateway.NewFileSystemWallet("./wallet")
	if err != nil {
		http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
		return
	}

	gw, err := gateway.Connect(
		wallet,
		gateway.WithConfig(gateway.FromFile("../gateway/connection-org1.yaml")),
		gateway.WithIdentity("appUser"),
	)
	if err != nil {
		http.Error(w, "Failed to connect to gateway", http.StatusInternalServerError)
		return
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		http.Error(w, "Failed to get network", http.StatusInternalServerError)
		return
	}

	contract := network.GetContract("carbon_contract")

	_, err = contract.SubmitTransaction("LogCarbonData", req.ShipmentData, "carbon_emission_data")
	if err != nil {
		http.Error(w, "Failed to submit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submitTransaction", submitTransaction).Methods("POST")

	http.ListenAndServe(":8080", r)
}