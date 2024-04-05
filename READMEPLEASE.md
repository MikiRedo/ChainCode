
Trabajamos en un solo canal que desplegaremos desde ./network.sh.
Desplegaremos un solo Chaincode


"Chaincode1" tiene dos funciones:

    1: Una función de registro donde el productor añade el objeto que crea.
    2: Una función donde el transportista y el distribuidor pueden actualizar el estado.

Creamos (primeramente) un Struct donde añadimos los atributos del objeto
Este "struct" nos servirá para que el productor añada las caracteristicas y los transportistas/distribuidores puedan actualizar el estado

