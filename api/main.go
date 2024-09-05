package main

import (
	"crypto/rand"
	"example/web-service-gin/docs"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Bank Mock Code of client Data //
type Client struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Surname       string  `json:"surname"`
	AccountNumber string  `json:"accountnumber"`
	BankID        string  `json:"bankid"`
	Balance       float64 `json:"balance"`
}
type Bank struct {
	BankID  string `json:"bankid"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

///////////////////////////////////

// send this to smart contract
type transaction struct {
	TransactionID   string         `json:"transaction_id"`
	ClientDetails   AccountDetails `json:"client_details"`
	ReceiverDetails AccountDetails `json:"receiver_details"`

	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

type AccountDetails struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	AccountNumber string `json:"accountnumber"`
	BankDetails   Bank   `json:"bankdetails"`
}

///////////////////////////////////

// payload used to start a new transaction
type payload struct {
	Amount          float64        `json:"amount"`
	ReceiverDetails AccountDetails `json:"receiverdetails"`
}

///////////////////////////////////

// MockData
var Sender = Client{
	ID:            "alexsid",
	Name:          "Alex",
	Surname:       "Dateling",
	AccountNumber: "0000000000",
	BankID:        "ABSA645334",
	Balance:       4000.00}

var SenderBank = Bank{
	BankID:  "ABSA645334",
	Name:    "ABSA",
	Country: "ZAR"}

///////////////////////////////////

/////////////////////////////////// FUNCTIONS

func generateTransactionReference(clientID string, receiverAccountNumber string) (string, error) {
	// Get the current timestamp in milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Generate a random number
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000)) // up to 999999
	if err != nil {
		return "", err
	}

	// Format the reference number as a string
	reference := fmt.Sprintf("%d%s%s%06d", timestamp, clientID, receiverAccountNumber, randomNumber.Int64())

	return reference, nil
}

// Function to check the balance of the client to see if there is enough funds to proceed with the transaction
func checkBalance(AccountHolder Client, amount float64) (result bool) {
	if AccountHolder.Balance < amount {
		fmt.Println("no enough munney")
		return false
	} else {
		return true
	}
}

// send transaction to chaincode
// func sendtoChaincode(newTransaction transaction) (result bool) {

// }

/////////////////////////////////// MOCK DATA FUNCTIONS

// Function to get the mockdata of the client
func getClient(clientId string) (result Client) {
	if clientId == Sender.ID {
		fmt.Println("yay")
		return Sender
	} else {
		fmt.Println("gay")
		return
	}
}

// Function to get the mockdata of the client's bank
func getBank(bankId string) (result Bank) {
	if bankId == SenderBank.BankID {
		fmt.Println("yay")
		return SenderBank
	} else {
		fmt.Println("gay")
		return
	}
}

/////////////////////////////////// API Calls

func createPayment(c *gin.Context) {
	clientid := c.Param("clientid")

	var newPayload payload

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	// Bind the received JSON to newTransaction
	if err := c.BindJSON(&newPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(newPayload)

	//validate payload

	// create transaction json from payload and client stuff

	var transaction_id string

	loginClient := getClient(clientid)
	clientBank := getBank(loginClient.BankID)

	fmt.Println(clientBank)

	// Need to generate a proper reference
	transaction_id, err := generateTransactionReference(loginClient.ID, newPayload.ReceiverDetails.AccountNumber)
	if err != nil {
		fmt.Println("Error generating reference number:", err)
		return
	}

	fmt.Println(transaction_id)

	var newTransaction transaction

	var isEnough bool = checkBalance(loginClient, newPayload.Amount)

	if isEnough {
		newTransaction.TransactionID = transaction_id
		newTransaction.ClientDetails = AccountDetails{
			Name:          loginClient.Name,
			Surname:       loginClient.Surname,
			AccountNumber: loginClient.AccountNumber,
			BankDetails:   clientBank,
		}
		newTransaction.ReceiverDetails = newPayload.ReceiverDetails
		newTransaction.Amount = newPayload.Amount
		newTransaction.Status = "initiated"

		fmt.Println("XXXXXXXXXXXXXX")
		fmt.Println(newTransaction)

		// sendtoChaincode()
		c.IndentedJSON(http.StatusCreated, newTransaction)
	} else {
		c.IndentedJSON(http.StatusBadRequest, `{"status": "failed"}`)
	}
}

// Send to chaincode
// func sendLedger(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // getAlbums responds with the list of all albums as JSON.
// func validateTransaction(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // getAlbums responds with the list of all albums as JSON.
// func postTransactionSmartContract(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

func main() {

	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.Default()
	router.POST("client/:clientid/createPayment", createPayment)
	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8080")
}

// workflow
// client A makes a payment
// bank A checks if funds is enough
// bank A reduces client balance
// bank A then proceeds after verification
// Sends it to smart contract (call it SC)
// SC verifies that bank A is legit
// SC makes sure that bank B is legit????
// SC creates a ledger thing...??
// ledger states that there is now a contract with Bank A and Bank B for X amount with a timestamped hash
// SC comes back saying everything is good
// SC gives a reference of contract to both Bank A and Bank B
// bank A now has a contract with bank B "stating hey we need to pay bank B"
// Bank B
