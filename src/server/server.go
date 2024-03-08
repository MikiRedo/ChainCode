package server

import (
	"core-functions-cc/config"
	"errors"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

var logger = flogging.MustGetLogger("chaincode-server")

// NewChaincodeServer devuelve una instancia de ChaincodeServer inicializada
func NewChaincodeServer(
	config *config.ChaincodeConfig,
	chaincode *contractapi.ContractChaincode,
) (chaincodeServerInstance *shim.ChaincodeServer, err error) {
	chaincodeID, err := getChaincodeID(config)
	if err != nil {
		return
	}

	chaincodeServerInstance = &shim.ChaincodeServer{
		CC:      chaincode,
		CCID:    chaincodeID,
		Address: config.ChaincodeAddress,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	logger.Info("Created chaincode server instance.")

	return
}

func getChaincodeID(config *config.ChaincodeConfig) (chaincodeID string, err error) {
	if config.ChaincodeID != "" {
		chaincodeID = config.ChaincodeID
		return
	}

	err = errors.New("CHAINCODE_ID env var must be set to start the server")
	return
}
