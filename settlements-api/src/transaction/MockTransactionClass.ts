import * as crypto from "crypto";
import { AccountDetails, Bank, hashedAccountDetails, Payload, Transaction } from 'src/models/transaction.model';

export class MockTransaction {
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

  private static bankCache: Map<string, Bank> = new Map();

  static createMockBank(): Bank {
    const sadcCountry = this.sadcCountries[Math.floor(Math.random() * this.sadcCountries.length)];
    const bankName = this.bankNames[Math.floor(Math.random() * this.bankNames.length)];
    const fullBankName = `${sadcCountry.country} ${bankName}`;
    
    // Check if this bank already exists in our cache
    const cacheKey = `${sadcCountry.country}-${bankName}`;
    if (this.bankCache.has(cacheKey)) {
      return this.bankCache.get(cacheKey)!;
    }

    // If not, create a new bank
    const countryCode = sadcCountry.country.toLowerCase().replace(/\s+/g, '');
    const bankNameCode = bankName.toLowerCase().replace(/\s+/g, '');
    const newBank: Bank = {
      bankid: `${countryCode}-${bankNameCode}`,
      name: fullBankName,
      country: sadcCountry.country
    };

    // Cache the new bank
    this.bankCache.set(cacheKey, newBank);
    return newBank;
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
    // HASHING CANT get the payload back
    // const jsonString = JSON.stringify(accountDetails);
    // const hash = crypto.createHash('sha256').update(jsonString).digest('hex');
    // return hash;



    const jsonString = JSON.stringify(accountDetails);

    const buffer = Buffer.from(jsonString, 'utf-8');

    // Convert Buffer to Base64 string
    const base64String = buffer.toString('base64');

    return base64String;
  }

  static decodeTransactionID(transaction_id: string): hashedAccountDetails {
    // Create a Buffer from the Base64 string
    const buffer = Buffer.from(transaction_id, 'base64');
  
    // Convert Buffer to JSON string
    const jsonString = buffer.toString('utf-8');
  
    // Parse JSON string into a JSON object
    return JSON.parse(jsonString);
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