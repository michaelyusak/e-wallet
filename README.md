# Sea Wallet

### check branch dev

Sea Wallet is a money management application. You can top up to your wallet, and transfer to ther Sea Wallet account. You also can customize you profile name and picture.

## API documentation
Documentation for this E Wallet REST API can be seen in
```
https://documenter.getpostman.com/view/33040861/2sA2rCTgtC
```

## Requirement
* Go
* Postgres
* node.js

## How to Run
* Create empty database
* Switch to the database and run ./backend/sql/ddl.sql and ./backend/sql/dml.sql, or use \i command in PSQL Shell
* Make sure you have created .env file
* In backend directory, run `go mod tidy`, then run `go run .`
* In frontend directory, run `npm i`, then `npm start`