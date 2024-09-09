import { Injectable } from '@nestjs/common';
import { Bank, Client, Payload, Transaction } from '../models/transaction.model';

@Injectable()
export class TransactionService {
    
    private transactions: Transaction[] = [];

    createTransaction(payload: Payload, clientid: string): Transaction {
        var newTransaction = {} as Transaction;
        newTransaction.transaction_id = Date.now().toString();

        // generate a transaction
        console.log("Getting Client details")
        var loginClient : Client
        loginClient = getClient(clientid)   // need to update this to get a client

        var clientBank : Bank
        console.log("Getting Clients Bank details")
	    clientBank = getBank(loginClient.bankid) // need to get client bank details

        // check if client has enough
        var isEnough : boolean = checkBalance(loginClient, payload.amount)

        if (isEnough){
            //   generate a new transactionID
            newTransaction.transaction_id = Date.now().toString();

            newTransaction.client_details = {
                name : loginClient.name,
                surname:       loginClient.surname,
                accountnumber: loginClient.accountnumber,
                bankdetails:   clientBank,
                
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
  
    settleTransactionPayment(TransactionID: string): Transaction | undefined {
      const transaction = this.getTransaction(TransactionID);
      if (transaction) {
        transaction.status = 'settled';
      }

    // call chaincode to settle transaction

    // update the mockdata transactions
      return transaction;
    }

    settleTransactionReceive(TransactionID: string): Transaction | undefined {
        const transaction = this.getTransaction(TransactionID);
        if (transaction) {
          transaction.status = 'settled';
        }
        return transaction;
      }
  
    getAllTransactions(bankid: string): Transaction[] {
      return this.transactions.filter(
        t => t.client_details.bankdetails.bankid === bankid || t.receiver_details.bankdetails.bankid === bankid
      );
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
var Sender : Client = {
	id:           "alexsid",
	name:          "Alex",
	surname:       "Dateling",
	accountnumber: "0000000000",
	bankid:        "ABSA645334",
	balance:       4000.00
}

var SenderBank : Bank = {
	bankid:  "ABSA645334",
	name:    "ABSA",
	country: "ZAR"}