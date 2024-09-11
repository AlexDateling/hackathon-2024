package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing an settlement
type SmartContract struct {
	contractapi.Contract
}

// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type PaymentDetails struct {
	PaymentID       string         `json:"payment_id"`
	ClientDetails   AccountDetails `json:"client_details"`
	ReceiverDetails AccountDetails `json:"receiver_details"`
	Amount          float64        `json:"amount"`
}
type Settlement struct {
	SettlementID        string `json:"settlement_id"`
	SenderBankDetails   Bank   `json:"sender_bank_details"`
	ReceiverBankDetails Bank   `json:"receiver_details"`

	Status         string `json:"status"`
	ClientStatus   string `json:"clientstatus"`
	ReceiverStatus string `json:"receiverstatus"`

	TimeEpoch string `json:"time_epoch"`
}

type AccountDetails struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	AccountNumber string `json:"accountnumber"`
	BankDetails   Bank   `json:"bankdetails"`
}

type Bank struct {
	BankID  string `json:"bankid"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// InitLedger adds a base set of settlements to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	settlements := []Settlement{}

	for _, settlement := range settlements {
		settlementJSON, err := json.Marshal(settlement)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(settlement.SettlementID, settlementJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) CreatePayment(ctx contractapi.TransactionContextInterface, paymentId string, clientdetails AccountDetails, receiverdetails AccountDetails, amount float64) error {
	// exists, err := s.PaymentExists(ctx, paymentId)
	// if err != nil {
	// 	return err
	// }
	// if exists {
	// 	return fmt.Errorf("the payment %s already exists", paymentId)
	// }
	// payment := PaymentDetails{
	// 	PaymentID:       paymentId,
	// 	ClientDetails:   clientdetails,
	// 	ReceiverDetails: receiverdetails,
	// 	Amount:          amount,
	// }

	// paymentJSON, err := json.Marshal(payment)
	// if err != nil {
	// 	return err
	// }

	// return ctx.GetStub().PutState(paymentId, paymentJSON)

	return nil
}

func (s *SmartContract) CreatePayment2(ctx contractapi.TransactionContextInterface, payment PaymentDetails) error {
	// exists, err := s.PaymentExists(ctx, payment.PaymentID)
	// if err != nil {
	// 	return err
	// }
	// if exists {

	// 	return fmt.Errorf("the settlement %s already exists", payment.PaymentID)
	// }

	// settlementJSON, err := json.Marshal(payment)
	// if err != nil {
	// 	return err
	// }

	// return ctx.GetStub().PutState(payment.PaymentID, settlementJSON)

	return nil
}

func (s *SmartContract) CreateSettlement(ctx contractapi.TransactionContextInterface, settlementId string, senderbankdetails Bank, receiverbankdetails Bank, status string, clientstatus string, receiverstatus string, timeepoch string) error {
	exists, err := s.SettlementExists(ctx, settlementId)
	if err != nil {
		return err
	}
	if exists {

		return fmt.Errorf("the settlement %s already exists", settlementId)
	}

	settlement := Settlement{
		SettlementID:        settlementId,
		SenderBankDetails:   senderbankdetails,
		ReceiverBankDetails: receiverbankdetails,
		Status:              status,
		ClientStatus:        clientstatus,
		ReceiverStatus:      receiverstatus,
		TimeEpoch:           timeepoch,
	}

	settlementJSON, err := json.Marshal(settlement)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(settlementId, settlementJSON)
}

func (s *SmartContract) CreateSettlement2(ctx contractapi.TransactionContextInterface, settlement Settlement) error {
	exists, err := s.SettlementExists(ctx, settlement.SettlementID)
	if err != nil {
		return err
	}
	if exists {

		return fmt.Errorf("the settlement %s already exists", settlement.SettlementID)
	}

	settlementJSON, err := json.Marshal(settlement)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(settlement.SettlementID, settlementJSON)
}

// ReadSettlement returns the settlement stored in the world state with given id.
func (s *SmartContract) ReadSettlement(ctx contractapi.TransactionContextInterface, settlementId string) (*Settlement, error) {
	settlementJSON, err := ctx.GetStub().GetState(settlementId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if settlementJSON == nil {
		return nil, fmt.Errorf("the settlement %s does not exist", settlementId)
	}

	var settlement Settlement
	err = json.Unmarshal(settlementJSON, &settlement)
	if err != nil {
		return nil, err
	}

	return &settlement, nil
}

// UpdateSettlement updates an existing settlement in the world state with provided parameters.
func (s *SmartContract) UpdateSettlement(ctx contractapi.TransactionContextInterface, settlementId string, status string) error {
	exists, err := s.SettlementExists(ctx, settlementId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the settlement %s does not exist", settlementId)
	}

	settlement, err := s.ReadSettlement(ctx, settlementId)
	if err != nil {
		return err
	}
	settlementJSON, err := json.Marshal(settlement)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(settlementId, settlementJSON)
}

// SettlementExists returns true when settlement with given ID exists in world state
func (s *SmartContract) SettlementExists(ctx contractapi.TransactionContextInterface, settlementId string) (bool, error) {
	settlementJSON, err := ctx.GetStub().GetState(settlementId)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return settlementJSON != nil, nil
}

// GetAllSettlements returns all settlement found in world state
func (s *SmartContract) GetAllSettlements(ctx contractapi.TransactionContextInterface) ([]*Settlement, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all settlements in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var settlements []*Settlement
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var settlement Settlement
		err = json.Unmarshal(queryResponse.Value, &settlement)
		if err != nil {
			return nil, err
		}
		settlements = append(settlements, &settlement)
	}

	return settlements, nil
}

// ADHOC Settlement
func (s *SmartContract) AdhocSettlePaymentSettlement(ctx contractapi.TransactionContextInterface, settlementId string) (*Settlement, error) {
	settlement, err := s.ReadSettlement(ctx, settlementId)
	if err != nil {
		return nil, err
	}

	settlement.ClientStatus = "SETTLED"

	// if both receiver and payment settlement statuses are setled settlement is regarded as completed.
	if settlement.ReceiverStatus == "SETTLED" && settlement.ClientStatus == "SETTLED" {
		settlement.Status = "COMPLETED"
	}

	settlementJSON, err := json.Marshal(settlement)
	if err != nil {
		return settlement, err
	}
	err = ctx.GetStub().PutState(settlement.SettlementID, settlementJSON)
	if err != nil {
		return settlement, err
	}
	// returns the settlement for the bank to do whatever they want with it
	return settlement, nil
}

func (s *SmartContract) AdhocSettleReceiveSettlement(ctx contractapi.TransactionContextInterface, settlementId string) (*Settlement, error) {
	settlement, err := s.ReadSettlement(ctx, settlementId)
	if err != nil {
		return nil, err
	}

	settlement.ReceiverStatus = "SETTLED"

	// if both receiver and payment settlement statuses are setled settlement is regarded as completed.
	if settlement.ReceiverStatus == "SETTLED" && settlement.ClientStatus == "SETTLED" {
		settlement.Status = "COMPLETED"
	}

	settlemenetJSON, err := json.Marshal(settlement)
	if err != nil {
		return settlement, err
	}
	err = ctx.GetStub().PutState(settlement.SettlementID, settlemenetJSON)
	if err != nil {
		return settlement, err
	}
	// returns the settlement for the bank to do whatever they want with it
	return settlement, nil
}
