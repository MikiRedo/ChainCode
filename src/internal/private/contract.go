package private

import (
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PrivateContract implementa las funciones privadas
type PrivateContract struct{}

// SetPrivateData guarda una clave en una colección privada de datos
func (c *PrivateContract) SetPrivateData(
	ctx contractapi.TransactionContextInterface,
	key string,
	value string,
	collection string,
) (result string, err error) {
	// Habitualmente, las transacciones con datos privadas deben hacerse usando transient y seleccionando endorsing peers:
	// transientMap, err := ctx.GetStub().GetTransient()
	// if err != nil {
	// 	err = errors.New("Error getting transient: " + err.Error())
	// 	return
	// }

	err = ctx.GetStub().PutPrivateData(collection, key, []byte(value))
	if err != nil {
		err = errors.New("Failed to set private data: " + err.Error())
		return
	}
	result = "{\"success\": true}"
	return
}

// GetPrivateData recupera una clave de una colección privada de datos
func (c *PrivateContract) GetPrivateData(
	ctx contractapi.TransactionContextInterface,
	key string,
	collection string,
) (data string, err error) {
	dataAsBytes, err := ctx.GetStub().GetPrivateData(collection, key)
	if err != nil {
		err = errors.New("Failed to get private data: " + err.Error())
	}

	data = string(dataAsBytes)
	return
}
