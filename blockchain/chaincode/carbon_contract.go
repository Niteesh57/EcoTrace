package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type CarbonData struct {
	ShipmentID string  `json:"shipmentID"`
	Emission   float64 `json:"emission"`
}

func (s *SmartContract) LogCarbonData(ctx contractapi.TransactionContextInterface, qrData string) error {
	var data CarbonData
	err := json.Unmarshal([]byte(qrData), &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	carbonJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(data.ShipmentID, carbonJSON)
}