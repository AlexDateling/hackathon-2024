package main

import (
	"example/web-service-gin/docs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Client struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	BankNumber string  `json:"banknumber"`
	BankID     string  `json:"bankid"`
	Balance    float64 `json:"balance"`
}

type Bank struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// send this to smart contract
type transaction struct {
	ID             string  `json:"id"`
	ClientSend     Client  `json:"clientsend"`
	ClientReceive  Client  `json:"clientreceive"`
	RequestingBank Bank    `json:"requestingbank"`
	BankOwed       Bank    `json:"bankowed"`
	AmountOwed     float64 `json:"amount"`
}

// albums slice to seed record album data.
var Sender = Client{
	ID:         "alexsid",
	Name:       "Alex",
	Surname:    "Dateling",
	BankNumber: "0000000000",
	BankID:     "ABSA645334",
	Balance:    4000.00}

var SenderBank = Bank{
	ID:      "ABSA645334",
	Name:    "ABSA",
	Country: "ZAR"}

var Receiver = Client{
	ID: "mumumumumumu",
	Name:       "T",
	Surname:    "Kok",
	BankNumber: "11111111111",
	BankID:     "FNB20003",
	Balance:    6.00}

// getAlbums responds with the list of all albums as JSON.
func createPayment(c *gin.Context) {
	clientid := c.Param("clientid")

	var newTransaction transaction

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	// Bind the received JSON to newTransaction
	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginClient := getClient(clientid)
	clientBank := getBank(loginClient.BankID)

	// fmt.Println(clientBank)

	newTransaction.ClientSend = loginClient
	newTransaction.RequestingBank = clientBank

	var isEnough bool = checkBalance(newTransaction.ClientSend, newTransaction.AmountOwed)

	// Add the new album to the slice.
	if isEnough {
		fmt.Println(newTransaction)
		c.IndentedJSON(http.StatusCreated, newTransaction)
	}

}

// getAlbums responds with the list of all albums as JSON.
func getClient(clientId string) (result Client) {
	if clientId == Sender.ID {
		fmt.Println("yay")
		return Sender
	} else {
		fmt.Println("gay")
		return
	}
}

func getBank(bankId string) (result Bank) {
	if bankId == SenderBank.ID {
		fmt.Println("yay")
		return SenderBank
	} else {
		fmt.Println("gay")
		return
	}
}

// getAlbums responds with the list of all albums as JSON.
func checkBalance(AccountHolder Client, amount float64) (result bool) {
	if AccountHolder.Balance < amount {
		fmt.Println("no enough munney")
		return false
	} else {
		return true
	}
}

// // getAlbums responds with the list of all albums as JSON.
// func getTimeStamp(c *gin.Context) {
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
