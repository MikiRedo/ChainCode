package config

import (
	"os"

	"github.com/hyperledger/fabric/common/flogging"
)

// ChaincodeConfig contiene la configuracion que leemos de variables de entorno
type ChaincodeConfig struct {
	ChaincodeAddress string
	ChaincodeID      string
}

// NewChaincodeConfig crea una nueva instancia de configuración y asigna las variables de entorno
func NewChaincodeConfig() *ChaincodeConfig {
	logger := flogging.MustGetLogger("chaincode-configuration")

	address, exists := os.LookupEnv("CHAINCODE_SERVER_ADDRESS")
	if !exists {
		address = "0.0.0.0:9999"
	} else {
		logger.Infof("CHAINCODE_SERVER_ADDRESS=%s", address)
	}

	chaincodeID, exists := os.LookupEnv("CHAINCODE_ID")
	if !exists {
		logger.Warn("CHAINCODE_ID environment variable is not set.")
	} else {
		logger.Infof("CHAINCODE_ID=%s", chaincodeID)
	}

	config := new(ChaincodeConfig)
	config.ChaincodeAddress = address
	config.ChaincodeID = chaincodeID

	return config
}
