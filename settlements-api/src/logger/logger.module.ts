import { Module } from '@nestjs/common';
import { ApplicationLogger } from './application.logger';

@Module({
  providers: [ApplicationLogger],
  exports: [ApplicationLogger],
})
export class LoggerModule {}
