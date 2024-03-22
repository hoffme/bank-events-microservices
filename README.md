# Modelado de Banco con microservicios orientados a eventos

## Endpoints
- [Postman](/docs/Bank.postman_collection.json)

## Teconologias
- **nginx**: proxy inverso como router y entrada principal
- **mongodb**: base de datos no-sql, misma instancia pero diferentes bbdd por microservicio
- **rabbitmq**: para la comunicacion por evento entre los microservicios
- **golang**: lenguaje con el que se desarrollo el servicio de transacciones
- **python**: lenguaje con el que se desarrollo el servicio de cuentas

## Como iniciar:
-   ir a la raiz del proyecto
-   ejecutar: `docker compose up -d --build`

## Modelado
- servicio backend-accounts:
    - **descripcion**:
        administracion de las cuentas
    - **acciones**
        - obtener todas las cuentas filtrando
        - obtener una cuenta con uuid
        - crear una cuenta con uuid, nombre, moneda y balance incial
        - modificar el nombre de la cuenta
        - desactivar una cuenta
        - activar una cuenta
    - **eventos**
        - actualiza los balances de las cuentas cuando se realiza una transaccion
- servicio backend-transacions:
    - **descripcion**:
        administrar las transacciones de las cuentas 
    - **acciones**
        - crear una transaccion
        - autorizar una transaccion mediante una cola
        - mantiene los balances de las cuentas correlacionados con las transacciones operadas
    - **eventos**
        - escucha cuando se crea o modifica una cuenta en el servicio de cuentas y la crea o modifica en la base de datos local
        - cuando se crea una transaccion se pone pendiente hasta que la cola de autorizacion la complete o rechaze, ahi se realiza el movimiento de balances en las cuentas

## Flujo:
- Se crea una cuenta en el servicio de cuentas en el endpoint `PUT /accounts/:account_id`
    -   parametros:
        -   id: uuid de la cuenta
        -   name: nombre de la cuenta
        -   currency: moneda de la cuenta
        -   balance: balance inicial de la cuenta

    -   acciones:
        -   guarda en la base de datos de cuentas
        -   dispara envento de cuenta creada a rabbitmq

    -   side effects:
        -   se crea una cuenta en el servicio de transacciones
        -   apartir de ahora el servicio de transacciones puede ver la cuenta

- Se crea una transacciones en el servicio de transacciones en el endpoint `PUT /transactions/:transaction_id`
    -   parametros:
        -   id: uuid de la transaccion
        -   from_account_id: id de la cuenta desde que se transfiere
        -   to_account_id: id de la cuenta hacia la que se transfiere
        -   currency: moneda de la transferencia
        -   amount: monto de la transferencia

    -   acciones:
        -   guarda en la base de datos de transacciones una transaccion pendiente
        -   dispara envento de transaccione creada a rabbitmq
        -   cola en el servicio de tranacciones se autoriza la tranaccion (para evitar condiciones de carrera)
        -   si se compelta la transaccion o se rechaza se emite un evento
    
    -   side effects
        -   otra cola en el servicio de transacciones actualiza la cuenta en de transacciones y lanza un evento de cuenta actualizada
        -   cola del servicio de cuentas escucha el evento y actualiza el balance de la cuenta
