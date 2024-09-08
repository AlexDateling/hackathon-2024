import { Module } from '@nestjs/common';
import { ContractService } from './contract.service';
import { LoggerModule } from '../logger/logger.module';

@Module({
  providers: [ContractService],
  exports: [ContractService],
  imports: [LoggerModule],
})
export class ContractModule {}
