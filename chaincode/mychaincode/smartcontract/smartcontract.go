package chaincode

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

type Transaction struct {
	TransactionID   string         `json:"transaction_id"`
	ClientDetails   AccountDetails `json:"client_details"`
	ReceiverDetails AccountDetails `json:"receiver_details"`

	Amount         float64 `json:"amount"`
	Status         string  `json:"status"`
	ClientStatus   string  `json:"clientstatus"`
	ReceiverStatus string  `json:"receiverstatus"`
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
		{TransactionID: "1725575573618alexsid9845543456653834",
			ClientDetails: AccountDetails{
				Name:          "Alex",
				Surname:       "Dateling",
				AccountNumber: "0000000000",
				BankDetails: Bank{
					BankID:  "ABSA645334",
					Name:    "ABSA",
					Country: "ZAR",
				},
			},
			ReceiverDetails: AccountDetails{
				Name:          "NotAlex1",
				Surname:       "ASurname123",
				AccountNumber: "9845543456",
				BankDetails: Bank{
					BankID:  "BANKID12345",
					Name:    "ABank",
					Country: "RSA",
				},
			},
			Amount:         2000,
			Status:         "PENDING",
			ClientStatus:   "PENDING",
			ReceiverStatus: "PENDING",
		},
		{TransactionID: "1723618alexsid984554343834",
			ClientDetails: AccountDetails{
				Name:          "T",
				Surname:       "T",
				AccountNumber: "11111111",
				BankDetails: Bank{
					BankID:  "ASDDSA",
					Name:    "ABSA",
					Country: "ZAR",
				},
			},
			ReceiverDetails: AccountDetails{
				Name:          "vvvv",
				Surname:       "assasaas",
				AccountNumber: "0877654",
				BankDetails: Bank{
					BankID:  "BANKID12345",
					Name:    "ABank",
					Country: "RSA",
				},
			},
			Amount:         600,
			Status:         "PENDING",
			ClientStatus:   "SETTLED",
			ReceiverStatus: "PENDING",
		},
		{
			TransactionID: "1723618alexsid984554343834",
			ClientDetails: AccountDetails{
				Name:          "T",
				Surname:       "T",
				AccountNumber: "11111111",
				BankDetails: Bank{
					BankID:  "ASDDSA",
					Name:    "ABSA",
					Country: "ZAR",
				},
			},
			ReceiverDetails: AccountDetails{
				Name:          "vvvv",
				Surname:       "assasaas",
				AccountNumber: "0877654",
				BankDetails: Bank{
					BankID:  "BANKID12345",
					Name:    "ABank",
					Country: "RSA",
				},
			},
			Amount:         600,
			Status:         "PENDING",
			ClientStatus:   "PENDING",
			ReceiverStatus: "SETTLED",
		},
		{
			TransactionID: "2837465alexsid984554343835",
			ClientDetails: AccountDetails{
				Name:          "Alice",
				Surname:       "Smith",
				AccountNumber: "22223333",
				BankDetails: Bank{
					BankID:  "BANKXYZ",
					Name:    "Bank of Example",
					Country: "USD",
				},
			},
			ReceiverDetails: AccountDetails{
				Name:          "Bob",
				Surname:       "Jones",
				AccountNumber: "7654321",
				BankDetails: Bank{
					BankID:  "BANKXYZ678",
					Name:    "Another Bank",
					Country: "US",
				},
			},
			Amount:         1500,
			Status:         "SETTLED",
			ClientStatus:   "SETTLED",
			ReceiverStatus: "SETTLED",
		},
		{
			TransactionID: "3948576alexsid984554343836",
			ClientDetails: AccountDetails{
				Name:          "John",
				Surname:       "Doe",
				AccountNumber: "33334444",
				BankDetails: Bank{
					BankID:  "XYZ123",
					Name:    "Global Bank",
					Country: "GBP",
				},
			},
			ReceiverDetails: AccountDetails{
				Name:          "Jane",
				Surname:       "Doe",
				AccountNumber: "43211234",
				BankDetails: Bank{
					BankID:  "XYZ456",
					Name:    "Regional Bank",
					Country: "GB",
				},
			},
			Amount:         2500,
			Status:         "SETTLED",
			ClientStatus:   "SETTLED",
			ReceiverStatus: "SETTLED",
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

func (s *SmartContract) CreateTransaction(ctx contractapi.TransactionContextInterface, transactionId string, clientDetails AccountDetails, receiverDetails AccountDetails, amount float64, status string, clientstatus string, receiverstatus string) error {
	exists, err := s.TransactionExists(ctx, transactionId)
	if err != nil {
		return err
	}
	if exists {

		return fmt.Errorf("the transaction %s already exists", transactionId)
	}

	transaction := Transaction{
		TransactionID:   transactionId,
		ClientDetails:   clientDetails,
		ReceiverDetails: receiverDetails,
		Amount:          amount,
		Status:          status,
		ClientStatus:    clientstatus,
		ReceiverStatus:  receiverstatus,
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

// Once a week batch call, bank will grab all payment transactions with their ID and then return those transactions
func (s *SmartContract) SettlePaymentTransactions(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {

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
			if transaction.ClientDetails.BankDetails.BankID == bankID {
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
func (s *SmartContract) SettleReceiveTransactions(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {

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
			if transaction.ReceiverDetails.BankDetails.BankID == bankID {
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
