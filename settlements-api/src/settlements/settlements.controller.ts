import { Body, Controller, Post, Req } from '@nestjs/common';
import { ContractService } from '../contract/contract.service';

@Controller('settlements')
export class SettlementsController {
  constructor(private readonly contractService: ContractService) {}

  @Post()
  async Post(@Req() req, @Body() settlementCreateRequest: any) {
    const params = [];

    params.push(JSON.stringify(settlementCreateRequest));

    await this.contractService.invokeChainCode(
      'CreateSettlementRequest',
      params,
      req.user.email,
    );

    return {}; // to decided the response
  }
}
