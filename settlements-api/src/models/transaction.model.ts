import { ApiProperty } from "@nestjs/swagger";  // helps Swagger generate accurate documentation for your API models.

export class Client {
  @ApiProperty()
  id: string;
  @ApiProperty()
  name: string;
  @ApiProperty()
  surname: string;
  @ApiProperty()
  accountnumber: string;
  @ApiProperty()
  bankid: string;
  @ApiProperty()
  balance: number;
}

export class Bank {
  @ApiProperty()
  bankid: string;
  @ApiProperty()
  name: string;
  @ApiProperty()
  country: string;
}

export class AccountDetails {
  @ApiProperty()
  name: string;
  @ApiProperty()
  surname: string;
  @ApiProperty()
  accountnumber: string;
  @ApiProperty()
  bankdetails: Bank;
}

export class Transaction {
  @ApiProperty()
  transaction_id: string;
  @ApiProperty()
  client_details: AccountDetails;
  @ApiProperty()
  receiver_details: AccountDetails;
  @ApiProperty()
  amount: number;
  @ApiProperty()
  status: string;
  @ApiProperty()
  clientstatus: string;
  @ApiProperty()
  receiverstatus: string;
    newTransaction: { name: string; surname: any; };
}

export class Payload {
  @ApiProperty()
  amount: number;
  @ApiProperty()
  receiverdetails: AccountDetails;
}