package public

import (
	"core-functions-cc/internal/models"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/flogging"
)

// PublicContract implementa las funciones públicas
type PublicContract struct{}

// Foo devuelve {'foo': 'bar'}
func (c *PublicContract) Foo() string {
	return "{\"foo\": \"bar\"}"
}

// Set actualiza una clave y devuelve {success: true}
func (c *PublicContract) Set(
	ctx contractapi.TransactionContextInterface,
	key string,
	value string,
) (result string, err error) {
	if internalError := ctx.GetStub().PutState(key, []byte(value)); internalError != nil {
		err = errors.New("Failed to store data: " + err.Error())
		return
	}

	result = "{\"success\": true}"
	return
}

// Get devuelve una clave del world state
func (c *PublicContract) Get(
	ctx contractapi.TransactionContextInterface,
	key string,
) (result string, err error) {
	data, internalError := ctx.GetStub().GetState(key)
	if internalError != nil {
		err = errors.New("Failed to get data: " + err.Error())
		return
	}
	if data == nil {
		err = errors.New("Data does not exist")
		return
	}

	result = fmt.Sprintf("\"%s\"", data)
	return
}

// CreateEvent emita y guarda un evento en el ledger
func (c *PublicContract) CreateEvent(
	ctx contractapi.TransactionContextInterface,
	key string,
	value string,
) (result string, err error) {
	event := models.Event{
		ID:      []string{"test-tx-id"},
		Type:    "test-tx-type",
		Payload: []byte("{\"notification\": {\"message\": \"hi\", \"type\": \"test-tx-type\", \"recipients\": []}, \"body\": " + "\"" + value + "\"}"),
	}

	// composite key para identificar el asset de forma unica.
	basicKey := append(event.ID, event.Type)
	compositeKey, err := ctx.GetStub().CreateCompositeKey(key, basicKey)
	if err != nil {
		flogging.MustGetLogger("core-functions-test").Error(err)
		return
	}

	var payload []byte
	// serializamos el evento
	if payload, err = json.Marshal(event); err != nil {
		flogging.MustGetLogger("core-functions-test").Error(err)
		return
	}

	if err = ctx.GetStub().SetEvent(compositeKey, payload); err != nil {
		flogging.MustGetLogger("core-functions-test").Error(err)
		return
	}

	if err = ctx.GetStub().PutState(compositeKey, payload); err != nil {
		flogging.MustGetLogger("core-functions-test").Error(err)
		return
	}

	result = fmt.Sprintf("\"%s\"", payload)
	return
}

// GetCreator deuvelve la identidad del creador de la transacción
func (c *PublicContract) GetCreator(
	ctx contractapi.TransactionContextInterface,
) (result string, err error) {
	creator, err := cid.GetMSPID(ctx.GetStub())
	if err != nil {
		err = errors.New("Failed to get creator: " + err.Error())
		return
	}

	result = "{\"creator\":" + "\"" + creator + "\"}"
	return
}

// CreateLog genera un log predeterminado
func (c *PublicContract) CreateLog() string {
	logger := flogging.MustGetLogger("GenerateLog")
	logger.Warningf("This log must not be included in the response.")

	return "{\"success\": true}"
}
