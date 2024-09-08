import { Injectable, Scope, HttpException, HttpStatus } from '@nestjs/common';
import { Wallets, Gateway, Wallet } from 'fabric-network';
import * as FabricCAServices from 'fabric-ca-client';
import * as fs from 'fs';
import * as path from 'path';
import { ApplicationLogger } from '../logger/application.logger';

@Injectable({ scope: Scope.REQUEST })
export class ContractService {
  private readonly ccpPath: string;
  private readonly wallet: Promise<Wallet>;

  constructor(private readonly logger: ApplicationLogger) {
    this.ccpPath = path.resolve(
      __dirname,
      '..',
      '..',
      'config',
      process.env.CONNECTION_PROFILE,
    );
    this.wallet = Wallets.newCouchDBWallet({ url: process.env.WALLET_URL });
  }

  private async getConnectionProfile() {
    const ccpJSON = await fs.promises.readFile(this.ccpPath, 'utf8');
    return JSON.parse(ccpJSON);
  }

  private async connectGateway(user: string) {
    const ccp = await this.getConnectionProfile();
    const userWallet = await this.getUserWallet(user);

    if (!userWallet) {
      throw new Error('User does not exist.');
    }

    const gateway = new Gateway();
    await gateway.connect(ccp, {
      wallet: userWallet,
      identity: user,
      discovery: { enabled: true },
    });

    return gateway;
  }

  private async getContract(gateway: Gateway) {
    const network = await gateway.getNetwork(process.env.APPLICATION_CHANNEL);
    return network.getContract(process.env.ORDER_CONTRACT_NAME);
  }

  async invokeChainCode(chaincode: string, payload: string[], user: string) {
    let gateway: Gateway;
    try {
      gateway = await this.connectGateway(user);
      const contract = await this.getContract(gateway);

      const result = await contract.submitTransaction(chaincode, ...payload);
      this.logger.log('Query result: ' + result);
      return result;
    } catch (error) {
      this.logger.error(error, 'invokeChainCode');
      throw new HttpException(
        error.toString(),
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    } finally {
      if (gateway) await gateway.disconnect();
    }
  }

  async queryChainCode(chaincode: string, params: string[], user: string) {
    let gateway: Gateway;
    try {
      gateway = await this.connectGateway(user);
      const contract = await this.getContract(gateway);

      const result = await contract.evaluateTransaction(chaincode, ...params);
      const stringResult = result.toString();

      if (!stringResult) {
        return { Key: 'No Records Found', Record: null };
      }

      return JSON.parse(stringResult);
    } catch (error) {
      this.logger.error(error, 'queryChainCode');
      throw new HttpException(
        error.toString(),
        HttpStatus.INTERNAL_SERVER_ERROR,
      );
    } finally {
      if (gateway) await gateway.disconnect();
    }
  }

  async getUserWallet(username: string) {
    const wallet = await this.wallet;
    const userExists = await wallet.get(username);
    if (!userExists) {
      this.logger.log(
        `An identity for ${username} does not exist in the wallet`,
      );
      return null;
    }
    return wallet;
  }

  async createUserWallet(username: string) {
    try {
      const ccp = await this.getConnectionProfile();
      const wallet = await this.wallet;
      const gateway = new Gateway();
      await gateway.connect(ccp, {
        wallet,
        identity: process.env.ORG_ADMIN,
        discovery: { enabled: true },
      });

      const adminIdentity = await wallet.get(username);
      if (!adminIdentity) {
        throw new Error(
          `An identity for the enroll user "${username}" does not exist in the wallet`,
        );
      }

      const caInfo = ccp.certificateAuthorities[process.env.CA_NAME];
      const ca = new FabricCAServices(caInfo.url);

      const provider = wallet
        .getProviderRegistry()
        .getProvider(adminIdentity.type);
      const adminUser = await provider.getUserContext(adminIdentity, 'enroll');

      const secret = await ca.register(
        {
          affiliation: 'NApex.settlements',
          enrollmentID: username,
          role: 'client',
        },
        adminUser,
      );

      const enrollment = await ca.enroll({
        enrollmentID: username,
        enrollmentSecret: secret,
      });

      const identity = {
        credentials: {
          certificate: enrollment.certificate,
          privateKey: enrollment.key.toBytes(),
        },
        mspId: process.env.MSPID,
        type: 'X.509',
      };
      await wallet.put(username, identity);
    } catch (error) {
      this.logger.error(error, 'createUserWallet');
      throw error;
    }
  }

  async enrollAdmin(username: string) {
    try {
      const ccp = await this.getConnectionProfile();
      const caURL =
        ccp.certificateAuthorities[process.env.CERTIFICATE_AUTHORITY].url;
      const ca = new FabricCAServices(caURL);
      const wallet = await this.wallet;

      const adminIdentity = await wallet.get(username);
      if (!adminIdentity) {
        throw new Error(
          `An identity for the enroll user ${username} does not exist in the wallet`,
        );
      }

      const enrollment = await ca.enroll({
        enrollmentID: process.env.ORG_ADMIN,
        enrollmentSecret: process.env.ORG_ADMIN_SECRET,
      });

      const identity = {
        credentials: {
          certificate: enrollment.certificate,
          privateKey: enrollment.key.toBytes(),
        },
        mspId: process.env.MSPID,
        type: 'X.509',
      };
      await wallet.put(username, identity);
    } catch (error) {
      this.logger.error(error, 'enrollAdmin');
      throw error;
    }
  }
}
