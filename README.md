# Go API Starter SQL

Note: This is just a simple starter project I created while learning golang. A lot of improvement can be done. Please don't use it without reviewing the code first

## Tech Stack
* Go
* MySql

## Feature 
* Dockerized
* Database using gorm
* Migration
* Config using dotenv
* Validation using validator
* Route using Mux
* Auth using JWT

### Installation
Make sure you have docker installed.

1. Clone the repo

    ```git clone git@github.com:dev-parvej/go-api-starter-sql.git```
2. Copy env from env.example

    ```cp .env.example .env```

3. Run

    ```docker-compose up --build```

You only have to include ```--build``` on the first run. 

You can also run the project without docker. For that you have to install the MySql and Go separately and then run 

```go main.go```
