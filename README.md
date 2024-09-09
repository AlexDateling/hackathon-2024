# NApex Settlements/Payment Fabric Network

## Overview

The SADC Settlements Network is an innovative financial infrastructure project designed to revolutionize cross-border payments and settlements among South African and SADC (Southern African Development Community) banks. This repository contains the implementation of a Hyperledger Fabric network, which serves as the backbone for this cutting-edge system.

The project is a collaborative effort involving key stakeholders such as the South African Reserve Bank (SARB), ABSA, and other major financial institutions in the SADC region. By leveraging the power of blockchain technology, specifically Hyperledger Fabric, this network aims to address long-standing challenges in cross-border transactions, including high costs, lengthy processing times, and lack of transparency.

Key components of this project include:

1. **Hyperledger Fabric Network**: A permissioned blockchain network that ensures secure, transparent, and efficient transaction processing.

2. **Smart Contracts (Chaincode)**: Custom-developed smart contracts that encode the business logic for cross-border settlements, compliance checks, and currency conversions.

3. **Settlements API**: A robust API built with NestJS that provides a user-friendly interface for banks to interact with the blockchain network, submit transactions, and query the ledger.

4. **Test Network**: A fully functional test environment that mimics the production setup, allowing for thorough testing and development of new features.

This project represents a significant step forward in regional financial integration, aligning with the SADC's goals of economic cooperation and development. By providing a more efficient, cost-effective, and transparent system for cross-border transactions, the SADC Settlements Network has the potential to boost trade, improve financial inclusion, and contribute to the overall economic growth of the region.

## Use Case Scenario

The SADC Settlements Network enables banks in South Africa and other SADC countries to:

1. Process cross-border payments quickly and securely
2. Perform real-time settlements between participating banks
3. Maintain a distributed ledger of all transactions for improved transparency and auditability

## Benefits

- **Reduced Transaction Time**: Faster settlement of cross-border transactions compared to traditional systems
- **Lower Costs**: Reduced fees associated with intermediary banks and currency conversion
- **Increased Transparency**: All participating banks have access to a shared, immutable ledger of transactions
- **Enhanced Security**: Utilizes Hyperledger Fabric's robust security features to protect sensitive financial data
- **Improved Liquidity Management**: Real-time settlement allows for better cash flow management
- **Regulatory Compliance**: Built-in compliance with SADC region financial regulations
- **Scalability**: Easily extendable to include more banks and financial institutions within the SADC region

## Repository Structure

```
/
├── fabric-network/          # Scripts for building and managing the Hyperledger Fabric test network
├── settlements-api/         # NestJS-based API for interacting with the settlement network
├── chaincode/               # Smart contracts (chaincode) for the Hyperledger Fabric network
└── README.md
```

### Directory Details

- [`/fabric-network`](/fabric-network): Contains scripts and configuration files for setting up and managing the Hyperledger Fabric test network. This includes node setup, channel creation, and chaincode deployment scripts.

- [`/settlements-api`](/settlements-api): A NestJS-based API that provides an interface for banks to interact with the settlement network. It handles user authentication, transaction submission, and query operations.

- [`/chaincode`](/chaincode): Contains the smart contracts (chaincode) that define the business logic for cross-border settlements on the Hyperledger Fabric network. This includes functions for initiating payments, updating balances, and querying transaction history.

## Getting Started

(Include instructions for setting up the development environment, running the test network, and deploying the API and chaincode)

## Contact

Proudly Brought to you by Team **NApex**
Contributors:

- Tristan Kok
- Neo Motsumi
- Mluleki Phemelo Ndala
- Alex Dateling
- Hannah Foster
