# hackathon-2024

https://github.com/HyperledgerHandsOn/trade-finance-logistics
https://docs.stitch.money/payment-products/settlements


Call api

localhost:8080/client/alexsid/createPayment
{
	"ID": "dsadsadsadsadsads",
	"ClientReceive": {
        "ID": "mumumumumumu",
	    "Name":    "T",
	    "Surname": "Kok",
        "BankNumber": "11111111111",
	    "BankID": "notabank20003",
	    "Balance": 6.00
    },
	"BankOwed": {
        "ID": "notabank20003",
        "Name": "notabank",
        "Country": "RSA"
    },
	"Amount": 2000
}


workflow::
    - Bank API will call a "create transaction"
        - this call requires the customer details who is logged in
            - accountNumber
                - can check all the information based on that customer
        - it will require the payload of the details of the transaction
            - who is the money going to
                - what details does this customer have
                - bank details
            - how much
    - the bank then verifies internally:
        - hey does this customer even have enough money for this?
        - is this customer legit
    - the bank then needs to prep an "Asset" for the smart Contract containing details for the ledger
        - the customer
        - the customer's bank
        - the Amount
        - the Hash to identify this transaction
        - the receiver
        - the receiver's bank
        - time it was created
        - state of the asset????
        
    - once the smart contract is created a reference must be returned or id of the asset
    - once in the world ledger the bank is then allowed to access the ledger
        - see details of the transaction
        -