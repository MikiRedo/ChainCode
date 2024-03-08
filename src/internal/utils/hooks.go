package utils

import (
	"errors"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TrackFunction utiliza time track para loguear el tiempo de ejecución de una función
func TrackFunction(
	ctx contractapi.TransactionContextInterface,
) (err error) {
	function, _ := ctx.GetStub().GetFunctionAndParameters()
	defer timeTrack(time.Now(), function)

	return
}

// UnknownFunction devuelve un error si la función no existe
func UnknownFunction(
	ctx contractapi.TransactionContextInterface,
) (err error) {
	return errors.New("Received unknown chaincode function name")
}
