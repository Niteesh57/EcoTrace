package main

import (
  "fmt"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
  contractapi.Contract
}

func (s *SmartContract) LogCarbonData(ctx contractapi.TransactionContextInterface, qrData string) error {
  // Validate and store QR data
  fmt.Printf("QR Data: %s", qrData)
  return ctx.GetStub().PutState(qrData, []byte(qrData))
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