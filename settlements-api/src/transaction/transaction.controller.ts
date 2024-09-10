import { Body, Controller, Get, Param, Post, Query } from '@nestjs/common';
import { TransactionService } from './transaction.service';
import { Payload, Transaction } from 'src/models/transaction.model';
import { ApiOperation, ApiParam, ApiQuery, ApiTags } from '@nestjs/swagger';

// to group related endpoints
@ApiTags('transactions')   
@Controller('transaction')
export class TransactionController {
     
    constructor(private readonly transactionService: TransactionService) {}

    @Post(':clientid/createPayment')
    @ApiOperation({ summary: 'Creates a new Transaction and adds it to the Hyperledger Fabric Network.' })
    @ApiParam({ name: 'clientid', type: 'string' })
    createTransaction(@Body() Payload: Payload, @Param('clientid') clientid: string) {
      return this.transactionService.createTransaction(Payload, clientid);
    }

    @Get(':transactionid/getTransaction')
    @ApiOperation({ summary: 'gets the Transaction details for the transactionid provided.' })
    @ApiParam({ name: 'transactionid', type: 'string' })
    getTransaction(@Param('transactionid') id: string) {
      return this.transactionService.getTransaction(id);
    }

    @Post(':transactionid/settlePayment')
    @ApiOperation({ summary: 'Adhoc Process to settle payment transaction, based on the transactionid.' })
    @ApiParam({ name: 'transactionid', type: 'string' })
    settleTransactionPayment(@Param('transactionid') TransactionID: string) {
      return this.transactionService.settleTransactionPayment(TransactionID);
    }

    @Post(':transactionid/settleReceive')
    @ApiOperation({ summary: 'Adhoc Process to settle receiving transaction, based on the transactionid.' })
    @ApiParam({ name: 'transactionid', type: 'string' })
    settleTransactionReceive(@Param('transactionid') TransactionID: string) {
      return this.transactionService.settleTransactionReceive(TransactionID);
    }
  
    @Get('/all')
    @ApiOperation({ summary: 'gets all transactions based on the bankID, for both payment and receiving sides' })
    @ApiQuery({ name: 'bankid', type: 'string' })
    getAllTransactions(@Query('bankid') bankid: string) {
      return this.transactionService.getAllTransactions(bankid);
    } 
}
