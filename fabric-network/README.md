# NApex Settlement Hyperlegder Fabric Network

This project re-establishes the Hyperledger [test-network](../test-network) as a _cloud native_ application.

## Prerequisites:

- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [jq](https://stedolan.github.io/jq/)
- [envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html) (`brew install gettext` on OSX)

- Kubernetes:
  - [KIND](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) + [Docker](https://www.docker.com) (resources: 8 CPU / 8 GRAM)

## Quickstart

Create a KIND cluster:

```shell
./network cluster init
```

Launch the network, create a channel, and deploy the [settlements-chaincode](../chaincode/) smart contract:

```shell
./network up

./network channel create

./network chaincode deploy settlements-chaincode ../settlements-chaincode/

```

Invoke and query chaincode:

```shell
./network chaincode invoke settlements-chaincode '{"Args":["InitLedger"]}'
./network chaincode query  settlements-chaincode '{"Args":["SettleReceiveTransactions","bankID"]}'
```
./network chaincode query  settlements-chaincode '{"Args":["chaincodefunction","json"]}'