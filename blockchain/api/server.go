package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		QRData string `json:"qrData"`
	}
	json.NewDecoder(r.Body).Decode(&requestData)

	wallet, err := gateway.NewFileSystemWallet(".wallet")
	if err != nil {
		http.Error(w, "Failed to create wallet", http.StatusInternalServerError)
		return
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile("../gateway/connection-org1.yaml"))
		gateway.WithIdentity(wallet, "appUser"))
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

	_, err = contract.SubmitTransaction("LogCarbonData", requestData.QRData)
	if err != nil {
		http.Error(w, "Failed to submit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submit", submitHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}