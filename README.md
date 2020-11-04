## Architecture
-  This is an attempt to implement a clean architecture, in case you don’t know it yet, here’s a reference https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
![Clean Architecture](clean.png)

## Objective

We have 2 types of users, common and merchant, both have wallet with money and make transfers between themselves.

Requirements:

- For both types of user, we need the Full Name, CPF, e-mail and Password. CPF/CNPJ and e-mails must be unique in the system. Therefore, your system should allow only one registration with the same CPF or e-mail address.

- Common users can send money (make transfers) to merchants and between common users.

- Merchant users only receive transfers, do not send money to anyone.

- Before finalizing the transfer, an external authorization service must be consulted, use this mock to simulate (https://run.mocky.io/v3/8fafdd68-a090-496f-8c9a-3442cf30dae6).

- The transfer operation must be a transaction (that is, reversed in any case of inconsistency) and the money must be returned to the sending user's wallet.

- Upon receiving payment, the common user or merchant needs to receive the notification sent by a third party service and eventually this service may become unavailable/unstable. Use this simulation to simulate sending (https://run.mocky.io/v3/b19f7b9f-9cbf-4fc6-ad22-dc30601aec04).

- This service must be RESTFul.

## Requirements/dependencies
    - Docker
    - Docker-compose

## Getting Started

- Starting API in port `:3001`

```sh
make start
```

- Run the tests using a container

```sh
make test
```

- Run the tests using a local machine

```sh
make test-local
```

- Run coverage

```sh
make coverage
```

- View the application logs

```sh
make logs
```

- Destroy application

```sh
make down
```

## API Endpoint

| Endpoint           | HTTP Method           | Description           |
| :----------------: | :-------------------: | :-------------------: |
| `/users`           | `POST`                | `Create user`         |
| `/users/{:userId}` | `GET`                 | `Find user by ID`     |
| `/transactions`    | `POST`                | `Create transaction`     |
| `/health`          | `GET`                 | `Health check`        |

## Test endpoints API using curl

- #### Creating new user

`Request`
```bash
curl --location --request POST 'localhost:3001/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fullname": "Gabriel Gabriel",
    "email": "gabriel@hotmail.com",
    "password": "passw123",
    "document": {
        "type": "CPF",
        "value": "070.910.549-64"
    },
    "wallet": {
        "currency": "BRL",
        "amount": 100
    },
    "type": "common"
}'
```

`Response`
```json
{
    "id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
    "full_name": "Common user",
    "email": "test@testing.com",
    "password": "passw",
    "document": {
        "type": "CPF",
        "value": "07091054954"
    },
    "wallet": {
        "currency": "BRL",
        "amount": 100
    },
    "roles": {
        "can_transfer": true
    },
    "type": "COMMON",
    "created_at": "0001-01-01T00:00:00Z"
}
```

- #### Find user by ID

`Request`
```bash
curl -i --request GET 'http://localhost:3001/users/{:userId}'
```

`Response`
```json
{
    "id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
    "fullname": "Common user",
    "email": "test@testing.com",
    "document": {
        "type": "CPF",
        "value": "07091054954"
    },
    "wallet": {
        "currency": "BRL",
        "amount": 100
    },
    "roles": {
        "can_transfer": true
    },
    "type": "COMMON",
    "created_at": "0001-01-01T00:00:00Z"
}
```

- #### Create new transaction

`Request`
```bash
curl --location --request POST 'localhost:3001/transactions' \
--header 'Content-Type: application/json' \
--data-raw '{
    "value": 100,
    "payer_id": {:userId},
    "payee_id": {:userId}
}'
```

`Response`
```json
{
    "id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
    "payer_id": "0db298eb-c8e7-4829-84b7-c1036b4f0791",
    "payee_id": "0db298eb-c8e7-4829-84b7-c1036b4f0792",
    "value": 100,
    "created_at": "0001-01-01T00:00:00Z"
}
```