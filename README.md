# Setup the application

## 1. Install the mysql
```bash
sudo apt update
sudo apt install mysql-server
sudo systemctl start mysql
sudo systemctl enable mysql
sudo mysql_secure_installation
sudo systemctl status mysql
sudo mysql -u root -p
```

## 2. Install the golang
```bash
wget https://golang.org/dl/go1.22.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.22.1.linux-amd64.tar.gz
nano ~/.bashrc

## Add following lines into .bashrc
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
### end

source ~/.bashrc
go version # check installation.
```


## 3. Fill your environment variable
.env file:
```bash
APP_PORT=9000

DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=amartha_billing_db # Change to anything you want
DB_USERNAME=root # Please change based on your mysql setup
DB_PASSWORD=root # Please change based on your mysql setup
```

## 4. Run the Makefile script
```bash
make db-create
make db-migrate
make db-seed
```

## 5. Run the golang app
```bash
make run
```


# API Documentation

This document describes the available API endpoints, their methods, and how to interact with them.

## Base URL
The base URL for the API is:
`http://<your-server>:<your_server_port>/api/`

## Endpoints

### 1. **Get Outstanding Balance**

- **Path**: `/outstanding`
- **Method**: `GET`
- **Description**: Fetch the outstanding balance for a loan or user.
- **Request**: This endpoint does not require any body parameters.
- **Response**: The response will contain the outstanding balance.

#### Example Request:
```bash
GET /api/outstanding?borrower_id=1
```

In Curl:
```bash
curl --location 'http://localhost:9000/api/outstanding?borrower_id=1'
```

#### Example Response:
```json
{
    "status": "success",
    "status_code": 200,
    "message": "success call api /api/outstanding",
    "metadata": null,
    "data": {
        "amount": 5500000
    }
}
```

### 2. **Is Delinquent**

- **Path**: `/is-delinquent`
- **Method**: `POST`
- **Description**: Fetch the check the delinquent status of the loan.
- **Request**: This endpoint does require json body parameters.
- **Response**: The response will contain the delinquent status in a boolean.

#### Example Request:
```bash
POST /api/is-delinquent
```

Headers:
- Content-Type: application/json

Body:
```json
{
    "borrower_id": 1
}
```

In Curl:
```bash
curl --location 'http://localhost:9000/api/is-delinquent' \
--header 'Content-Type: application/json' \
--data '{
    "borrower_id": 1
}'
```

#### Example Response:
```json
{
    "status": "success",
    "status_code": 200,
    "message": "success call api /api/is-delinquent",
    "metadata": null,
    "data": {
        "isDelinquent": false
    }
}
```

### 2. **Make Payment**

- **Path**: `/payment`
- **Method**: `POST`
- **Description**: Pay the pending payment at the current time.
- **Request**: This endpoint does require json body parameters.
- **Response**: The response will contain the nb of successful payment.

#### Example Request:
```bash
POST /api/payment
```

Headers:
- Content-Type: application/json

Body:
```json
{
    "loan_id": 1,
    "paid_amount": 300000
}
```

In Curl:
```bash
curl --location 'http://localhost:9000/api/payment' \
--header 'Content-Type: application/json' \
--data '{
    "loan_id": 1,
    "paid_amount": 300000
}'
```

#### Example Response:
```json
{
    "status": "success",
    "status_code": 200,
    "message": "success call api /api/payment",
    "metadata": null,
    "data": {
        "successfulPayments": 3
    }
}
```