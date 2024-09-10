package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing an Transaction
type SmartContract struct {
	contractapi.Contract
}

// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type hashedAccountDetails struct {
	ClientDetails   AccountDetails `json:"client_details"`
	ReceiverDetails AccountDetails `json:"receiver_details"`
	Amount          float64
	TimeEpoch       string
}
type Transaction struct {
	TransactionID       string `json:"transaction_id"`
	SenderBankDetails   Bank   `json:"sender_bank_details"`
	ReceiverBankDetails Bank   `json:"receiver_details"`

	Status         string `json:"status"`
	ClientStatus   string `json:"clientstatus"`
	ReceiverStatus string `json:"receiverstatus"`
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

// InitLedger adds a base set of transactions to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	transactions := []Transaction{

		{
			TransactionID: "07edce096e2e50b4ce06f9e2c21a5f2e0d62824b40d9ff57ac520ad865e419d2",
			SenderBankDetails: Bank{
				BankID:  "mauritius-nationalbank",
				Name:    "Mauritius National Bank",
				Country: "Mauritius",
			},
			ReceiverBankDetails: Bank{
				BankID:  "eswatini-agriculturalbank",
				Name:    "Eswatini Agricultural Bank",
				Country: "Eswatini",
			},
			Status:         "SETTLED",
			ClientStatus:   "PENDING",
			ReceiverStatus: "FAILED",
		},
		{
			TransactionID: "d2c0e825e3f41ce0d612e9cfd4b9efa74491326f497b9fda94598327641f2c28",
			SenderBankDetails: Bank{
				BankID:  "namibia-nationalbank",
				Name:    "Namibia National Bank",
				Country: "Namibia",
			},
			ReceiverBankDetails: Bank{
				BankID:  "tanzania-developmentbank",
				Name:    "Tanzania Development Bank",
				Country: "Tanzania",
			},
			Status:         "SETTLED",
			ClientStatus:   "PENDING",
			ReceiverStatus: "FAILED",
		},
		{
			TransactionID: "6ab502e0cdeae5bbcbeb6a1a5adc02c117d2fb7531df38306b44bada65b07418",
			SenderBankDetails: Bank{
				BankID:  "angola-nationalbank",
				Name:    "Angola National Bank",
				Country: "Angola",
			},
			ReceiverBankDetails: Bank{
				BankID:  "mozambique-investmentbank",
				Name:    "Mozambique Investment Bank",
				Country: "Mozambique",
			},
			Status:         "SETTLED",
			ClientStatus:   "FAILED",
			ReceiverStatus: "FAILED",
		},
		{
			TransactionID: "349c8379423d6e557fba23d0c71c08e6b23e681f332b95b1f5d90eb3b2e008e3",
			SenderBankDetails: Bank{
				BankID:  "seychelles-merchantbank",
				Name:    "Seychelles Merchant Bank",
				Country: "Seychelles",
			},
			ReceiverBankDetails: Bank{
				BankID:  "mozambique-cooperativebank",
				Name:    "Mozambique Cooperative Bank",
				Country: "Mozambique",
			},
			Status:         "FAILED",
			ClientStatus:   "FAILED",
			ReceiverStatus: "PENDING",
		},
		{
			TransactionID: "08427ac3c611563fea3f119858e85f3eb538813352396940b61fd8e11e9ccc7a",
			SenderBankDetails: Bank{
				BankID:  "zimbabwe-nationalbank",
				Name:    "Zimbabwe National Bank",
				Country: "Zimbabwe",
			},
			ReceiverBankDetails: Bank{
				BankID:  "tanzania-nationalbank",
				Name:    "Tanzania National Bank",
				Country: "Tanzania",
			},
			Status:         "PENDING",
			ClientStatus:   "FAILED",
			ReceiverStatus: "SETTLED",
		},
		{
			TransactionID: "025d70bffd79271dee158ac23d16fb8b06a5af44cc57a2bae1ae6a94a16337db",
			SenderBankDetails: Bank{
				BankID:  "namibia-investmentbank",
				Name:    "Namibia Investment Bank",
				Country: "Namibia",
			},
			ReceiverBankDetails: Bank{
				BankID:  "mauritius-agriculturalbank",
				Name:    "Mauritius Agricultural Bank",
				Country: "Mauritius",
			},
			Status:         "FAILED",
			ClientStatus:   "SETTLED",
			ReceiverStatus: "PENDING",
		},
		{
			TransactionID: "1bb1c832a9ce5ab3f6ca6a68571f0107c7e08c3f946115023c049997b87772a1",
			SenderBankDetails: Bank{
				BankID:  "botswana-agriculturalbank",
				Name:    "Botswana Agricultural Bank",
				Country: "Botswana",
			},
			ReceiverBankDetails: Bank{
				BankID:  "southafrica-merchantbank",
				Name:    "South Africa Merchant Bank",
				Country: "South Africa",
			},
			Status:         "SETTLED",
			ClientStatus:   "PENDING",
			ReceiverStatus: "SETTLED",
		},
		{
			TransactionID: "b86e8df8976c3034055f5ddd559c2173f0fff53c3f6a4a5b0d2f2fc6cd1fa668",
			SenderBankDetails: Bank{
				BankID:  "zambia-nationalbank",
				Name:    "Zambia National Bank",
				Country: "Zambia",
			},
			ReceiverBankDetails: Bank{
				BankID:  "zambia-merchantbank",
				Name:    "Zambia Merchant Bank",
				Country: "Zambia",
			},
			Status:         "SETTLED",
			ClientStatus:   "FAILED",
			ReceiverStatus: "FAILED",
		},
		{
			TransactionID: "f7e72980ee6437090e35b6c4e8aafaac0ea399c77450b9ec7b28a5e88a811733",
			SenderBankDetails: Bank{
				BankID:  "seychelles-merchantbank",
				Name:    "Seychelles Merchant Bank",
				Country: "Seychelles",
			},
			ReceiverBankDetails: Bank{
				BankID:  "democraticrepublicofcongo-commercialbank",
				Name:    "Democratic Republic of Congo Commercial Bank",
				Country: "Democratic Republic of Congo",
			},
			Status:         "FAILED",
			ClientStatus:   "PENDING",
			ReceiverStatus: "PENDING",
		},
		{
			TransactionID: "2effaf346fb267dd4a30782259dddd194b2e2b03cd469a24dffe1a2353e5d331",
			SenderBankDetails: Bank{
				BankID:  "comoros-savingsbank",
				Name:    "Comoros Savings Bank",
				Country: "Comoros",
			},
			ReceiverBankDetails: Bank{
				BankID:  "tanzania-commercialbank",
				Name:    "Tanzania Commercial Bank",
				Country: "Tanzania",
			},
			Status:         "FAILED",
			ClientStatus:   "FAILED",
			ReceiverStatus: "FAILED",
		},
	}

	for _, transaction := range transactions {
		transactionJSON, err := json.Marshal(transaction)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

func (s *SmartContract) CreateTransaction(ctx contractapi.TransactionContextInterface, transactionId string, senderbankdetails Bank, receiverbankdetails Bank, status string, clientstatus string, receiverstatus string) error {
	exists, err := s.TransactionExists(ctx, transactionId)
	if err != nil {
		return err
	}
	if exists {

		return fmt.Errorf("the transaction %s already exists", transactionId)
	}

	transaction := Transaction{
		TransactionID:       transactionId,
		SenderBankDetails:   senderbankdetails,
		ReceiverBankDetails: receiverbankdetails,
		Status:              status,
		ClientStatus:        clientstatus,
		ReceiverStatus:      receiverstatus,
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(transactionId, transactionJSON)
}

func (s *SmartContract) CreateTransaction2(ctx contractapi.TransactionContextInterface, transaction Transaction) error {
	exists, err := s.TransactionExists(ctx, transaction.TransactionID)
	if err != nil {
		return err
	}
	if exists {

		return fmt.Errorf("the transaction %s already exists", transaction.TransactionID)
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
}

// ReadTransaction returns the transaction stored in the world state with given id.
func (s *SmartContract) ReadTransaction(ctx contractapi.TransactionContextInterface, transactionId string) (*Transaction, error) {
	transactionJSON, err := ctx.GetStub().GetState(transactionId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if transactionJSON == nil {
		return nil, fmt.Errorf("the transaction %s does not exist", transactionId)
	}

	var transaction Transaction
	err = json.Unmarshal(transactionJSON, &transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// UpdateAsUpdateTransactionset updates an existing transaction in the world state with provided parameters.
func (s *SmartContract) UpdateTransaction(ctx contractapi.TransactionContextInterface, transactionId string, status string) error {
	exists, err := s.TransactionExists(ctx, transactionId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the transaction %s does not exist", transactionId)
	}

	transaction, err := s.ReadTransaction(ctx, transactionId)
	if err != nil {
		return err
	}
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(transactionId, transactionJSON)
}

// DO WE NEED THIS XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// // DeleteAsset deletes an given asset from the world state.
// func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
// 	exists, err := s.AssetExists(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return fmt.Errorf("the asset %s does not exist", id)
// 	}

// 	return ctx.GetStub().DelState(id)
// }

// TransactionExists returns true when transaction with given ID exists in world state
func (s *SmartContract) TransactionExists(ctx contractapi.TransactionContextInterface, transactionId string) (bool, error) {
	transactionJSON, err := ctx.GetStub().GetState(transactionId)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return transactionJSON != nil, nil
}

// GetAllTransactions returns all transactions found in world state
func (s *SmartContract) GetAllTransactions(ctx contractapi.TransactionContextInterface) ([]*Transaction, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all transactions in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transaction
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

// ADHOC Settlement
func (s *SmartContract) AdhocSettlePaymentTransactions(ctx contractapi.TransactionContextInterface, transactionId string) (*Transaction, error) {
	transaction, err := s.ReadTransaction(ctx, transactionId)
	if err != nil {
		return nil, err
	}

	transaction.ClientStatus = "SETTLED"

	// if both receiver and payment transaction statuses are setled transaction is regarded as completed.
	if transaction.ReceiverStatus == "SETTLED" && transaction.ClientStatus == "SETTLED" {
		transaction.Status = "COMPLETED"
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return transaction, err
	}
	err = ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
	if err != nil {
		return transaction, err
	}
	// returns the transactions for the bank to do whatever they want with it
	return transaction, nil
}

func (s *SmartContract) AdhocSettleReceiveTransactions(ctx contractapi.TransactionContextInterface, transactionId string) (*Transaction, error) {
	transaction, err := s.ReadTransaction(ctx, transactionId)
	if err != nil {
		return nil, err
	}

	transaction.ReceiverStatus = "SETTLED"

	// if both receiver and payment transaction statuses are setled transaction is regarded as completed.
	if transaction.ReceiverStatus == "SETTLED" && transaction.ClientStatus == "SETTLED" {
		transaction.Status = "COMPLETED"
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return transaction, err
	}
	err = ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
	if err != nil {
		return transaction, err
	}
	// returns the transactions for the bank to do whatever they want with it
	return transaction, nil
}

// Once a week batch call, bank will grab all payment transactions with their ID and then return those transactions
func (s *SmartContract) BatchSettlePaymentTransactions(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transaction
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}
		// Get all pending transactions
		if transaction.Status == "PENDING" {
			// Check which transaction is the bank that called it, need some sort of authorization aswell
			if transaction.SenderBankDetails.BankID == bankID {
				transactions = append(transactions, &transaction)
			}
		}

		// UPDATE CLIENT DETAILS OF BANK THAT CALLED API TO "SETTLED"
		// PLEASE CHECK IF THIS IS BYREF
		for _, transaction := range transactions {
			transaction.ClientStatus = "SETTLED"

			// if both receiver and payment transaction statuses are setled transaction is regarded as completed.
			if transaction.ReceiverStatus == "SETTLED" && transaction.ClientStatus == "SETTLED" {
				transaction.Status = "COMPLETED"
			}

			transactionJSON, err := json.Marshal(transaction)
			if err != nil {
				return transactions, err
			}
			err = ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
			if err != nil {
				return transactions, err
			}
		}
	}

	// returns the transactions for the bank to do whatever they want with it
	return transactions, nil
}

// Once a week batch call, bank will grab all transactions with their ID and then return those transactions
func (s *SmartContract) BatchSettleReceiveTransactions(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var transactions []*Transaction
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var transaction Transaction
		err = json.Unmarshal(queryResponse.Value, &transaction)
		if err != nil {
			return nil, err
		}
		// Get all pending transactions
		if transaction.Status == "PENDING" || transaction.Status == "INPROGRESS" {
			// Check which transaction is the bank that called it, need some sort of authorization aswell
			if transaction.ReceiverBankDetails.BankID == bankID {
				transactions = append(transactions, &transaction)
			}
		}

		// UPDATE CLIENT DETAILS OF BANK THAT CALLED API TO "SETTLED"
		// PLEASE CHECK IF THIS IS BYREF
		for _, transaction := range transactions {
			transaction.ReceiverStatus = "SETTLED"

			// if both receiver and payment transaction statuses are setled transaction is regarded as completed.
			if transaction.ReceiverStatus == "SETTLED" && transaction.ClientStatus == "SETTLED" {
				transaction.Status = "COMPLETED"
			} else {
				transaction.Status = "INPROGRESS"
			}

			transactionJSON, err := json.Marshal(transaction)
			if err != nil {
				return transactions, err
			}
			err = ctx.GetStub().PutState(transaction.TransactionID, transactionJSON)
			if err != nil {
				return transactions, err
			}
		}
	}

	return transactions, err
}
