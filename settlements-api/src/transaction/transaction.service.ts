import { Injectable, NotFoundException } from '@nestjs/common';
import { Bank, Client, Payload, Transaction } from '../models/transaction.model';

@Injectable()
export class TransactionService {

    private transactions: Transaction[] = [
        {
            transaction_id: "1725575573618alexsid9845543456653834",
            client_details: {
                name: "Alex",
                surname: "Dateling",
                accountnumber: "0000000000",
                bankdetails: {
                    bankid: "ABSA645334",
                    name: "ABSA",
                    country: "ZAR",
                },
            },
            receiver_details: {
                name: "NotAlex1",
                surname: "ASurname123",
                accountnumber: "9845543456",
                bankdetails: {
                    bankid: "BANKID12345",
                    name: "ABank",
                    country: "RSA",
                },
            },
            amount: 2000,
            status: "PENDING",
            clientstatus: "PENDING",
            receiverstatus: "PENDING",
        },
        {
            transaction_id: "1723618alexsid984554343834",
            client_details: {
                name: "T",
                surname: "T",
                accountnumber: "11111111",
                bankdetails: {
                    bankid: "ASDDSA",
                    name: "ABSA",
                    country: "ZAR",
                },
            },
            receiver_details: {
                name: "vvvv",
                surname: "assasaas",
                accountnumber: "0877654",
                bankdetails: {
                    bankid: "BANKID12345",
                    name: "ABank",
                    country: "RSA",
                },
            },
            amount: 600,
            status: "PENDING",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING",
        },
        {
            transaction_id: "1723618alexsid984554343834",
            client_details: {
                name: "T",
                surname: "T",
                accountnumber: "11111111",
                bankdetails: {
                    bankid: "ASDDSA",
                    name: "ABSA",
                    country: "ZAR",
                },
            },
            receiver_details: {
                name: "vvvv",
                surname: "assasaas",
                accountnumber: "0877654",
                bankdetails: {
                    bankid: "BANKID12345",
                    name: "ABank",
                    country: "RSA",
                },
            },
            amount: 600,
            status: "PENDING",
            clientstatus: "PENDING",
            receiverstatus: "SETTLED",
        },
        {
            transaction_id: "2837465alexsid984554343835",
            client_details: {
                name: "Alice",
                surname: "Smith",
                accountnumber: "22223333",
                bankdetails: {
                    bankid: "BANKXYZ",
                    name: "Bank of Example",
                    country: "USD",
                },
            },
            receiver_details: {
                name: "Bob",
                surname: "Jones",
                accountnumber: "7654321",
                bankdetails: {
                    bankid: "BANKXYZ678",
                    name: "Another Bank",
                    country: "US",
                },
            },
            amount: 1500,
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED",
        },
        {
            transaction_id: "3948576alexsid984554343836",
            client_details: {
                name: "John",
                surname: "Doe",
                accountnumber: "33334444",
                bankdetails: {
                    bankid: "XYZ123",
                    name: "Global Bank",
                    country: "GBP",
                },
            },
            receiver_details: {
                name: "Jane",
                surname: "Doe",
                accountnumber: "43211234",
                bankdetails: {
                    bankid: "XYZ456",
                    name: "Regional Bank",
                    country: "GB",
                },
            },
            amount: 2500,
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED",
        }
    ];

    createTransaction(payload: Payload, clientid: string): Transaction {
        var newTransaction = {} as Transaction;
        newTransaction.transaction_id = Date.now().toString();

        // generate a transaction
        console.log("Getting Client details")
        var loginClient: Client
        loginClient = getClient(clientid)   // need to update this to get a client

        var clientBank: Bank
        console.log("Getting Clients Bank details")
        clientBank = getBank(loginClient.bankid) // need to get client bank details

        // check if client has enough munney
        var isEnough: boolean = checkBalance(loginClient, payload.amount)

        if (isEnough) {
            //   generate a new transactionID
            newTransaction.transaction_id = Date.now().toString();

            newTransaction.client_details = {
                name: loginClient.name,
                surname: loginClient.surname,
                accountnumber: loginClient.accountnumber,
                bankdetails: clientBank,

            }
            newTransaction.receiver_details = payload.receiverdetails
            newTransaction.amount = payload.amount
            newTransaction.status = "PENDING"
            newTransaction.clientstatus = "PENDING"
            newTransaction.receiverstatus = "PENDING"

            console.log(newTransaction)

            // Call chaincode to push to ledger


            // throw it into array i guess
            this.transactions.push(newTransaction);
            return newTransaction;
        }
    }

    getTransaction(id: string): Transaction | undefined {
        // CALL chaincode to get the transaction

        // get it from some mockdata populate transactions with mockdata

        return this.transactions.find(t => t.transaction_id === id);
    }


    updateTransaction(TransactionID: string, updatedTransaction: Transaction): Transaction | undefined {
        const index = this.transactions.findIndex(t => t.transaction_id === TransactionID);
        if (index === -1) {
            throw new NotFoundException(`Transaction with ID ${TransactionID} not found`);
        }
        
        console.log(updatedTransaction)
        // Ensure the ID doesn't change
        updatedTransaction.transaction_id = TransactionID;
        
        // Replace the old transaction with the updated one
        this.transactions[index] = updatedTransaction;
        
        return this.transactions[index];
    }

    settleTransactionPayment(TransactionID: string): Transaction | undefined {
        const transaction = this.getTransaction(TransactionID);
        if (transaction) {
            transaction.clientstatus = 'SETTLED';
            if (transaction.clientstatus == 'SETTLED' && transaction.receiverstatus == 'SETTLED'){
                transaction.status = 'SETTLED'
            }
        }

        return this.updateTransaction(TransactionID, transaction)
        // call chaincode to settle transaction

        // update the mockdata transactions
        //  transaction;
    }

    settleTransactionReceive(TransactionID: string): Transaction | undefined {
        const transaction = this.getTransaction(TransactionID);
        if (transaction) {
            transaction.receiverstatus = 'SETTLED';
            if (transaction.clientstatus == 'SETTLED' && transaction.receiverstatus == 'SETTLED'){
                transaction.status = 'SETTLED'
            }
        }

        // call chaincode to settle transaction

        // update the mockdata transactions
        return this.updateTransaction(TransactionID, transaction)
    }

    getAllTransactions(bankid: string): Transaction[] {
        console.log("im in");
        var transactionsArr = this.transactions.filter(
            t => t.client_details.bankdetails.bankid === bankid || t.receiver_details.bankdetails.bankid === bankid
        )
        console.log(transactionsArr);
        return (transactionsArr);
    }

}

function getClient(clientid: string): Client {
    // Check MOCKDATA
    if (clientid == Sender.id) {
        console.log("yay there is a client with this id")
        return Sender
    } else {
        console.log("no user found")
        return
    }
}

function getBank(BankID: string): Bank {
    // Check MOCKDATA
    if (BankID == SenderBank.bankid) {
        console.log("found bank details with this ID")
        return SenderBank
    } else {
        console.log("no bank found with this id")
        return
    }
}

function checkBalance(loginClient: Client, amount: number) {
    if (loginClient.balance < amount) {
        console.log("not enough munney")
        return false
    } else {
        return true
    }
}

// MockData
var Sender: Client = {
    id: "alexsid",
    name: "Alex",
    surname: "Dateling",
    accountnumber: "0000000000",
    bankid: "ABSA645334",
    balance: 4000.00
}

var SenderBank: Bank = {
    bankid: "ABSA645334",
    name: "ABSA",
    country: "ZAR"
}