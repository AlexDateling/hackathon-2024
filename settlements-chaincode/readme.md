Building Chaincode
Now let’s compile your chaincode.

GO111MODULE=on go get -u github.com/hyperledger/fabric/core/chaincode/shim@v1.4
go build
Assuming there are no errors, now we can proceed to the next step, testing your chaincode.



Testing Using dev mode
Normally chaincodes are started and maintained by peer. However in “dev mode”, chaincode is built and started by the user. This mode is useful during chaincode development phase for rapid code/build/run/debug cycle turnaround.

We start “dev mode” by leveraging pre-generated orderer and channel artifacts for a sample dev network. As such, the user can immediately jump into the process of compiling chaincode and driving calls.


TRANSACTION STATUSES

PENDING
INPROGRESS
COMPLETED
INVALID???



NEED TO DO:
Ensure you have a connection profile for your Fabric network (usually a YAML file).
You'll need to have a wallet with the necessary identities to interact with the network.


Ref
https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html
https://github.com/hyperledger/fabric-samples

https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic


https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-private-data