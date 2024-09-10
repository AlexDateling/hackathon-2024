import { Injectable, NotFoundException } from '@nestjs/common';
import { Bank, Client, hashedAccountDetails, Payload, Transaction } from '../models/transaction.model';
import * as crypto from "crypto";
import { MockTransaction } from './MockTransactionClass';
@Injectable()
export class TransactionService {

    private transactions: Transaction[] = [
        {
            transaction_id: "f352a6557960e50caa15c6e1b34deb2ae303fc0cf59c0660c1f9508fb7b56ef0",
            Senderbankdetails: {
              bankid: "ZMW1",
              name: "Zambia Development Bank",
              country: "Zambia"
            },
            ReceiverBankDetails: {
              bankid: "ZMW2",
              name: "Zambia National Bank",
              country: "Zambia"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "f352a6557960e50caa15c6e1b34deb2ae303fc0cf59c0660c1f9508fb7b56ef0",
            Senderbankdetails: {
              bankid: "ZMW1",
              name: "Zambia Development Bank",
              country: "Zambia"
            },
            ReceiverBankDetails: {
              bankid: "ZMW2",
              name: "Zambia National Bank",
              country: "Zambia"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "f352a6557960e50caa15c6e1b34deb2ae303fc0cf59c0660c1f9508fb7b56ef0",
            Senderbankdetails: {
              bankid: "ZMW1",
              name: "Zambia Development Bank",
              country: "Zambia"
            },
            ReceiverBankDetails: {
              bankid: "ZMW2",
              name: "Zambia National Bank",
              country: "Zambia"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
        {
            transaction_id: "f352a6557960e50caa15c6e1b34deb2ae303fc0cf59c0660c1f9508fb7b56ef0",
            Senderbankdetails: {
              bankid: "ZMW1",
              name: "Zambia Development Bank",
              country: "Zambia"
            },
            ReceiverBankDetails: {
              bankid: "ZMW2",
              name: "Zambia National Bank",
              country: "Zambia"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "ac2c46fb1fd1ba3bc234b451364b0bf099373538c99fe85bfb3b2b7b5cc0448b",
            Senderbankdetails: {
              bankid: "MWK3",
              name: "Malawi National Bank",
              country: "Malawi"
            },
            ReceiverBankDetails: {
              bankid: "MGA4",
              name: "Madagascar Agricultural Bank",
              country: "Madagascar"
            },
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "FAILED"
          },
          {
            transaction_id: "76b87239f27f5d96266b7ee6617298a3d8e5a2185b312b4ff9ca2cdc971a3e65",
            Senderbankdetails: {
              bankid: "MGA5",
              name: "Madagascar Investment Bank",
              country: "Madagascar"
            },
            ReceiverBankDetails: {
              bankid: "KMF6",
              name: "Comoros Commercial Bank",
              country: "Comoros"
            },
            status: "PENDING",
            clientstatus: "FAILED",
            receiverstatus: "SETTLED"
          },
          {
            transaction_id: "376236860c0011302b4264df66a18bd394698446f51b4f2f8a831257b75d11cd",
            Senderbankdetails: {
              bankid: "MZN7",
              name: "Mozambique Cooperative Bank",
              country: "Mozambique"
            },
            ReceiverBankDetails: {
              bankid: "MWK8",
              name: "Malawi Development Bank",
              country: "Malawi"
            },
            status: "PENDING",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED"
          },
          {
            transaction_id: "16826996622510c8118ece9ff30ffb70bb3d7b7ffb515aaf0a073c7ab2e258be",
            Senderbankdetails: {
              bankid: "ZWL9",
              name: "Zimbabwe Development Bank",
              country: "Zimbabwe"
            },
            ReceiverBankDetails: {
              bankid: "SCR10",
              name: "Seychelles Commercial Bank",
              country: "Seychelles"
            },
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "ed205986106dc9ba91724297afde91018d408acf13ea44a35644a51625a6fd4e",
            Senderbankdetails: {
              bankid: "ZAR11",
              name: "South Africa Savings Bank",
              country: "South Africa"
            },
            ReceiverBankDetails: {
              bankid: "BWP12",
              name: "Botswana Agricultural Bank",
              country: "Botswana"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED"
          },
          {
            transaction_id: "3c73f1f55ca2357204d7fa59c48fff4b63de9f4d7dd6bb85cd12fc679f217796",
            Senderbankdetails: {
              bankid: "NAD13",
              name: "Namibia National Bank",
              country: "Namibia"
            },
            ReceiverBankDetails: {
              bankid: "SZL14",
              name: "Eswatini Cooperative Bank",
              country: "Eswatini"
            },
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED"
          },
          {
            transaction_id: "35529d190a78302220159f242df3fc41ae2ad68b47567a34b41b1d130824cffc",
            Senderbankdetails: {
              bankid: "BWP15",
              name: "Botswana Savings Bank",
              country: "Botswana"
            },
            ReceiverBankDetails: {
              bankid: "CDF16",
              name: "Democratic Republic of Congo Cooperative Bank",
              country: "Democratic Republic of Congo"
            },
            status: "FAILED",
            clientstatus: "FAILED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "480e71d8b9dcc9825b656633d491c6800a8640c9e1fdb67706774cb2b95ee064",
            Senderbankdetails: {
              bankid: "BWP17",
              name: "Botswana Agricultural Bank",
              country: "Botswana"
            },
            ReceiverBankDetails: {
              bankid: "LSL18",
              name: "Lesotho Agricultural Bank",
              country: "Lesotho"
            },
            status: "SETTLED",
            clientstatus: "FAILED",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "bf630dfcf38e3085a5b7c8cf29c675022d3c5870cebd3996ab7ac280bac3818c",
            Senderbankdetails: {
              bankid: "KMF19",
              name: "Comoros Cooperative Bank",
              country: "Comoros"
            },
            ReceiverBankDetails: {
              bankid: "ZAR20",
              name: "South Africa Development Bank",
              country: "South Africa"
            },
            status: "FAILED",
            clientstatus: "PENDING",
            receiverstatus: "FAILED"
          },
          {
            transaction_id: "d1a7e0452c54494cc4c2330ca1720fba53db9fa445ecf79e2b78f5df34cc7ffb",
            Senderbankdetails: {
              bankid: "SZL31",
              name: "Eswatini Development Bank",
              country: "Eswatini"
            },
            ReceiverBankDetails: {
              bankid: "ZMW1",
              name: "Zambia Development Bank",
              country: "Zambia"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          
          {
            transaction_id: "82e83bdd491ea40e73de549a8a47cf0658d378464f6a1dd7c013fbe6c38736fe",
            Senderbankdetails: {
              bankid: "ZWL33",
              name: "Zimbabwe Merchant Bank",
              country: "Zimbabwe"
            },
            ReceiverBankDetails: {
              bankid: "AOA34",
              name: "Angola Commercial Bank",
              country: "Angola"
            },
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
          },
          
          {
            transaction_id: "20e625fc7be29b1a28ec5cb77658859eba3698e4b951682ddd96cc3af27e4f12",
            Senderbankdetails: {
              bankid: "LSL35",
              name: "Lesotho Commercial Bank",
              country: "Lesotho"
            },
            ReceiverBankDetails: {
              bankid: "MUR36",
              name: "Mauritius Savings Bank",
              country: "Mauritius"
            },
            status: "PENDING",
            clientstatus: "PENDING",
            receiverstatus: "PENDING"
          },
          
          {
            transaction_id: "e321c95209209e7470b63d570f5ce0fd3b03031596555438f4442c82216aa95d",
            Senderbankdetails: {
              bankid: "CDF37",
              name: "Democratic Republic of Congo Agricultural Bank",
              country: "Democratic Republic of Congo"
            },
            ReceiverBankDetails: {
              bankid: "TZS38",
              name: "Tanzania Agricultural Bank",
              country: "Tanzania"
            },
            status: "SETTLED",
            clientstatus: "PENDING",
            receiverstatus: "SETTLED"
          },
          
          {
            transaction_id: "1ec99b623eadbb2fdf64dec4652a72d19a62769fb78d40b38c1d20b56bb8f3e2",
            Senderbankdetails: {
              bankid: "ZMW39",
              name: "Zambia Investment Bank",
              country: "Zambia"
            },
            ReceiverBankDetails: {
              bankid: "MZN40",
              name: "Mozambique Development Bank",
              country: "Mozambique"
            },
            status: "SETTLED",
            clientstatus: "PENDING",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "e872a6c98f9aa0e3e4be3ff6ce15416cb18c4038626c04fdecb8cf94beada902",
            Senderbankdetails: {
              bankid: "KMF41",
              name: "Comoros Investment Bank",
              country: "Comoros"
            },
            ReceiverBankDetails: {
              bankid: "BWP42",
              name: "Botswana Savings Bank",
              country: "Botswana"
            },
            status: "SETTLED",
            clientstatus: "PENDING",
            receiverstatus: "PENDING"
          },
          {
            transaction_id: "df0ee918dcdcb74070be58f625ba47d5e9d4883f78b6a8e02f99c50d4eb46728",
            Senderbankdetails: {
              bankid: "LSL43",
              name: "Lesotho Cooperative Bank",
              country: "Lesotho"
            },
            ReceiverBankDetails: {
              bankid: "MZN44",
              name: "Mozambique Savings Bank",
              country: "Mozambique"
            },
            status: "PENDING",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED"
          },
          
          {
            transaction_id: "3a1e3af4c7056de08cc75f9ad4bab25060374a1c8ef4e0fad4fb6f48db378821",
            Senderbankdetails: {
              bankid: "LSL45",
              name: "Lesotho Savings Bank",
              country: "Lesotho"
            },
            ReceiverBankDetails: {
              bankid: "ZWL46",
              name: "Zimbabwe Investment Bank",
              country: "Zimbabwe"
            },
            status: "SETTLED",
            clientstatus: "SETTLED",
            receiverstatus: "SETTLED"
          },
          
          {
            transaction_id: "024116b8701e1100cfa7a356b04cddb5a42c6796f058ffbe134b20b7878c0506",
            Senderbankdetails: {
              bankid: "KMF6",
              name: "Comoros Commercial Bank",
              country: "Comoros"
            },
            ReceiverBankDetails: {
              bankid: "MGA5",
              name: "Madagascar Investment Bank",
              country: "Madagascar"
            },
            status: "FAILED",
            clientstatus: "PENDING",
            receiverstatus: "SETTLED"
          },
          
          {
            transaction_id: "a669eaba64cf7e45665fc3ac4d4f18caf70ec11e43c5f7441587365876141abc",
            Senderbankdetails: {
              bankid: "SCR49",
              name: "Seychelles Savings Bank",
              country: "Seychelles"
            },
            ReceiverBankDetails: {
              bankid: "MWK50",
              name: "Malawi Cooperative Bank",
              country: "Malawi"
            },
            status: "FAILED",
            clientstatus: "SETTLED",
            receiverstatus: "PENDING"
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
        var isEnough: boolean = checkBalance()

        if (isEnough) {
            // XXXXXXXXXXXXXXXXXXXXXX HASHED
            var HashedSettlementDetails : hashedAccountDetails;
            HashedSettlementDetails.client_details = {
                name: loginClient.name,
                surname: loginClient.surname,
                accountnumber: loginClient.accountnumber,
                bankdetails: clientBank,

            }
            HashedSettlementDetails.receiver_details = payload.receiverdetails
            HashedSettlementDetails.amount = payload.amount
            HashedSettlementDetails.time_epoch = Date.now().toString()
            // XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

            // Details for the settlement/Transaction
            //  GENERATE HASH
            newTransaction.transaction_id = generateHashedDetails(HashedSettlementDetails)
            newTransaction.Senderbankdetails = clientBank
            newTransaction.ReceiverBankDetails = {
                bankid: payload.receiverdetails.bankdetails.bankid,
                name: payload.receiverdetails.bankdetails.name,
                country: payload.receiverdetails.bankdetails.country,
            }
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
            t => t.Senderbankdetails.bankid === bankid || t.ReceiverBankDetails.bankid === bankid
        )
        console.log(transactionsArr);
        return (transactionsArr);
    }

    generatemock() {

        // Generate and display 10 mock transactions
        console.log("Generating 10 mock transactions:");
        for (let i = 0; i < 10; i++) {
        console.log(`\nTransaction ${i + 1}:`);
        console.log(JSON.stringify(MockTransaction.createMockTransaction(), null, 2));
        }

        // Generate and display 10 mock payloads
        console.log("\nGenerating 10 mock payloads:");
        for (let i = 0; i < 10; i++) {
        console.log(`\nPayload ${i + 1}:`);
        console.log(JSON.stringify(MockTransaction.createMockPayload(), null, 2));
        }
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

function checkBalance() {
    return true
}

function generateHashedDetails(AccountDetails : hashedAccountDetails): string {
    const jsonString = JSON.stringify(AccountDetails);

    // Create a SHA-256 hash of the JSON string
    const hash = crypto.createHash('sha256');
    hash.update(jsonString);
    return hash.digest('hex');
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