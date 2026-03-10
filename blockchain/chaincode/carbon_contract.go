package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Shipment struct {
	ID      string `json:"id"`
	CarbonData string `json:"carbonData"`
}

func (s *SmartContract) LogCarbonData(ctx contractapi.TransactionContextInterface, shipmentID string, carbonData string) error {
	shipment := Shipment{
		ID: shipmentID,
		CarbonData: carbonData,
	}

	shipmentJSON, err := json.Marshal(shipment)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(shipmentID, shipmentJSON)
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}