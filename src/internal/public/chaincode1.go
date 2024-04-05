package public

import (
    "encoding/json"
    "fmt"
	"time"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"

)

//estructura del producto (no puedo añadirña en el repo "models")
type Product struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    Status     string `json:"status"` //("Registered", "Transported", "Delivered")
    Owner      string `json:"owner"`  //identificacion del propietario
    OwnerType  string `json:"ownerType"`  //("Fabricante", "Transportista", "Distribuidor")
    UpdateTime string `json:"updateTime"` //cuando se creo o actualizó
}

//el fabricante añade un nuevo producto
func NewProduct(ctx contractapi.TransactionContextInterface, id, name, owner string, ownerType string) error {
    product := Product{
        ID:        id,
        Name:      name,
        Status:    "Registered",
        Owner:     owner,     
        OwnerType: ownerType,
    }
    productBytes, err := json.Marshal(product)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(id, productBytes)  //no sabia como hacerlo 
}

//funcion para actualizar el estado del producto
func UpdateData(ctx contractapi.TransactionContextInterface, id, newStatus, owner string, ownerType string) error {
    productBytes, err := ctx.GetStub().GetState(id) //id del producto
    if err != nil {
        return fmt.Errorf("failed to read from world state: %v", err)
    }
    if productBytes == nil {
        return fmt.Errorf("this product %s does not exist", id)
    }

    var product Product
    err = json.Unmarshal(productBytes, &product)
    if err != nil {
        return err
    }

    //check que el owner i el ownerType que actualizan el producto sean de la empresa
	//si uno no cuadra, falla
    if product.Owner != owner || product.OwnerType != ownerType {
        return fmt.Errorf("only owner %s of type %s can update the status of product %s", owner, ownerType, id)
    }

    product.Status = newStatus
    product.Owner = owner
    product.OwnerType = ownerType
	// Obtener la marca de tiempo de la transacción (no sabia como hace el timestamp)
	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("error obteniendo la marca de tiempo de la transacción: %v", err)
	}
	txTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))

	//Formatear la marca de tiempo como una cadena en el formato RFC3339
	product.UpdateTime = txTime.Format(time.RFC3339)

    productBytes, err = json.Marshal(product)
    if err != nil {
        return err
    }
    return ctx.GetStub().PutState(id, productBytes)
}