package main

import (
	"core-functions-cc/internal/private"
	"core-functions-cc/internal/public"
	"core-functions-cc/internal/utils"

	"core-functions-cc/config"
	"core-functions-cc/server"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

func main() {
	chaincodeConfig := config.NewChaincodeConfig()

	smartContracts := new(SmartContractComposition)

	// Hooks para loggear el tiempo de ejecución - rechazar funciones incorrectas
	smartContracts.Contract.BeforeTransaction = utils.TrackFunction
	smartContracts.Contract.UnknownTransaction = utils.UnknownFunction

	chaincode, err := contractapi.NewChaincode(smartContracts)
	if err != nil {
		panic(err.Error())
	}

	chaincodeServerInstance, err := server.NewChaincodeServer(chaincodeConfig, chaincode)
	if err != nil {
		panic(err.Error())
	}

	flogging.
		MustGetLogger("main").
		Info("Started chaincode server.")

	if err := chaincodeServerInstance.Start(); err != nil {
		panic(err.Error())
	}
}

// SmartContractComposition combina todos los smart contracts en un struct
type SmartContractComposition struct {
	contractapi.Contract

	public.PublicContract
	private.PrivateContract
}
