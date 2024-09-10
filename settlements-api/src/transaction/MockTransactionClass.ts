import { Bank, AccountDetails, Transaction, hashedAccountDetails, Payload } from '../models/transaction.model'; // Assume these are imported from your models file
import crypto from 'crypto';

class MockTransaction {
  private static bankCounter = 1;
  private static accountCounter = 1;
  private static transactionCounter = 1;

  // SADC countries and their currencies
  private static sadcCountries = [
    { country: 'Angola', currency: 'AOA' },
    { country: 'Botswana', currency: 'BWP' },
    { country: 'Comoros', currency: 'KMF' },
    { country: 'Democratic Republic of Congo', currency: 'CDF' },
    { country: 'Eswatini', currency: 'SZL' },
    { country: 'Lesotho', currency: 'LSL' },
    { country: 'Madagascar', currency: 'MGA' },
    { country: 'Malawi', currency: 'MWK' },
    { country: 'Mauritius', currency: 'MUR' },
    { country: 'Mozambique', currency: 'MZN' },
    { country: 'Namibia', currency: 'NAD' },
    { country: 'Seychelles', currency: 'SCR' },
    { country: 'South Africa', currency: 'ZAR' },
    { country: 'Tanzania', currency: 'TZS' },
    { country: 'Zambia', currency: 'ZMW' },
    { country: 'Zimbabwe', currency: 'ZWL' }
  ];

  private static bankNames = [
    'National Bank', 'Commercial Bank', 'Savings Bank', 'Investment Bank',
    'Cooperative Bank', 'Development Bank', 'Agricultural Bank', 'Merchant Bank'
  ];

  static createMockBank(): Bank {
    const sadcCountry = this.sadcCountries[Math.floor(Math.random() * this.sadcCountries.length)];
    const bankName = this.bankNames[Math.floor(Math.random() * this.bankNames.length)];
    return {
      bankid: `${sadcCountry.currency}${this.bankCounter++}`,
      name: `${sadcCountry.country} ${bankName}`,
      country: sadcCountry.country
    };
  }

  static createMockAccountDetails(): AccountDetails {
    return {
      name: `User${this.accountCounter}`,
      surname: `Surname${this.accountCounter}`,
      accountnumber: `ACC${this.accountCounter++}`,
      bankdetails: this.createMockBank()
    };
  }

  static createMockHashedAccountDetails(): hashedAccountDetails {
    return {
      client_details: this.createMockAccountDetails(),
      receiver_details: this.createMockAccountDetails(),
      amount: Math.floor(Math.random() * 10000),
      time_epoch: Date.now().toString()
    };
  }

  static generateHashedDetails(accountDetails: hashedAccountDetails): string {
    const jsonString = JSON.stringify(accountDetails);
    const hash = crypto.createHash('sha256');
    hash.update(jsonString);
    return hash.digest('hex');
  }

  static createMockTransaction(): Transaction {
    const hashedDetails = this.createMockHashedAccountDetails();
    return {
      transaction_id: this.generateHashedDetails(hashedDetails),
      Senderbankdetails: hashedDetails.client_details.bankdetails,
      ReceiverBankDetails: hashedDetails.receiver_details.bankdetails,
      status: ['PENDING', 'SETTLED', 'FAILED'][Math.floor(Math.random() * 3)],
      clientstatus: ['PENDING', 'SETTLED', 'FAILED'][Math.floor(Math.random() * 3)],
      receiverstatus: ['PENDING', 'SETTLED', 'FAILED'][Math.floor(Math.random() * 3)]
    };
  }

  static createMockPayload(): Payload {
    return {
      amount: Math.floor(Math.random() * 10000),
      receiverdetails: this.createMockAccountDetails()
    };
  }
}

// Example usage
console.log(MockTransaction.createMockTransaction());
console.log(MockTransaction.createMockPayload());