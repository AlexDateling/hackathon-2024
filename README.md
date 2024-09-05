# hackathon-2024

https://github.com/HyperledgerHandsOn/trade-finance-logistics
https://docs.stitch.money/payment-products/settlements


Call api

localhost:8080/client/:id/createPayment

localhost:8080/client/alexsid/createPayment
ABSA Payload needed:
{
    Amount: X,
    ReceiverDetails: {
        Name: M
        Surname: S
        AccountNumber:1234567890
        ReceiverBankDetails: {
            BankID: notabank20003 <will link to the bank details>
            Name:
            Country: "RSA"
        }
    }
}

Function for "createPayment"
    - get the senders details
    - verify that the amount in the payload is less than the amount in the senders balance.





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
        - is this customer/receiver legit???????

    - the bank then needs to prep an "Asset" for the smart Contract containing details for the ledger

    TransactionDetails:
    {
        transaction_id
        sender_Details {

            Bank_Details: {

            }
        }
        
        receiver_details: {

            bank_details: {

            }
        }
        Amount:
        time it was created:
        
        state_of_the_asset:c    compliant_failed, payment_needed, Completed, 
    }

    compliant: {
        transaction_id:
        state_of_the_asset:???????
        valid
    }
        
    - once the smart contract is created a transaction_id must be returned or id of the asset
    - once in the world ledger the bank is then allowed to access the ledger
        - see details of the transaction



ToDo: 
    - how do we create this unique transaction_id
        - timestamped hash
    - how to send it to chaincode?
    - 
