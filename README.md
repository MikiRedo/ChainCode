# Core functions chaincode - GOLANG

## Introducción

Este chaincode ha sido desarrollado para el módulo **Cómo Desplegar una red de Fabric** del bootcamp de Blockchain de **Keepcoding**. Contiene una serie de funciones para explorar cómo funcionan las API para chaincodes en Golang, y que nos servirán para probar ciertas características de la plataforma en relación a smart contracts.

Incluye Dockerfile para ser desplegado como un servicio externo. 

## Uso

1. Despliega la test-network de forma habitual con couchDB
```
# Desde el directorio fabric-samples/test-network
./network.sh down
./network.sh up createChannel -s couchdb
```

1. Instala el chaincode como servicio externo junto con la definición de colleciones privadas de datos
```
./network.sh deployCCAAS  -ccn core -ccp ../../core-chaincode -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -cccg ../../core-chaincode/collections.json
```

1. Comprueba que ambos contenedores Docker han sido desplegados y están corriendo
```
docker ps -a | grep core
```

1. El chaincode está preparado para recibir transacciones

## Referencia de la interfaz

El chaincode contiene dos smart contracts (public & private) que contienen las siguientes funciones:
- Foo
- Set
- Get
- CreateEvent
- CreateLog
- GetCreator

- SetPrivateData
- GetPrivateData
