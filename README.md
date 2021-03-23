## Objective

Your assignment is to build an internal API for a fake financial institution using Go and any framework.

## Brief

While modern banks have evolved to serve a plethora of functions, at their core, banks must provide certain basic features. Today, your task is to build the basic HTTP API for one of those banks! Imagine you are designing a backend API for bank employees. It could ultimately be consumed by multiple frontends (web, iOS, Android etc).

## Data Model

The Banking API has 4 objects that make up the data model

### Customer

The customers of the Bank. Customers have the following attributes:

| Name | Description                        |
| ---- | ---------------------------------- |
| ID   | Primary identifier of the Customer |
| Name | Customers Name                     |

### Account Type

Lists the type of accounts that the Bank has. Account Types have the following attributes:

| Name | Description                            |
| ---- | -------------------------------------- |
| ID   | Primary identifier of the account type |
| Type | Name of the Account Type               |

### Account

The customers bank accounts at the Bank. Accounts have the following attributes:

| Name     | Description                                                    |
| -------- | -------------------------------------------------------------- |
| ID       | Primary identifier of the Account                              |
| Customer | Identifier of the Customer to whom this account belongs        |
| Type     | An identifier for the type of account. See Account Type above. |
| Balance  | The current balance for the account                            |

### Transaction

All the transactions that have been performed on a specific bank account. At the moment there is only one type of transaction, a transfer but with the addition of a Type attribute, this could be easily expanded to other transactions. Transactions have the following attributes:

| Name        | Description                                                          |
| ----------- | -------------------------------------------------------------------- |
| ID          | Primary identifier of the Transaction                                |
| CreatedAt   | Timestamp when the transaction took place                            |
| FromAccount | Identifier of the Bank Account when the transaction originated from  |
| ToAccount   | Identifier of the Bank Account when the transaction is being sent to |
| Amount      | Amount that the transaction was for                                  |

## Repositories

There are 3 repositories, one for each of the main data objects. These repositories would contain the bulk of the business logic that would be performed on each object

### Customers

The Customers repository has 2 main functions:

| Name         | Description                                                           |
| ------------ | --------------------------------------------------------------------- |
| GetCustomers | Returns all customers at the bank                                     |
| GetCustomer  | Receives the Identifier of the customer and returns a single customer |

### Accounts

The Accounts repository has 2 main functions:

| Name                  | Description                                                                                                                                                                                                           |
| --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreateCustomerAccount | Receives an new account object that contains all the accountattributes other than the Id. It returns the newly created account and it's new Identifier. An opening balance can be pass in as part of that new account |
| GetAccount            | Receives the Identifier of an account and returns that account object                                                                                                                                                 |
| GetCustomerAccount    | Receives both the Identifier of an account and the Identifier of a customer and returns that account object                                                                                                           |
| GetCustomerAccounts   | Receives a customer identifier and returns all accounts for that customer.                                                                                                                                            |

### Transactions

| Name                     | Description                                                                                                                                                                                                                                                                                                                      |
| ------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GetAccountTransactions   | receives both a Customer and Account Identifier and returns all transactions on that account. The function checks that the account identifer belongs to the customer identifier passed to the function. It returns transactions where it is From this account (returns the amount as a debit) or where it is To this account.    |
| CreateAccountTransaction | receives both a Customer and a From Account Identifier and a new transaction object and returns the newly created transaction. The function checks that the From account identifer belongs to the customer identifier passed to the function. It also checks that the To account identifier in the new tranaction object exists. |

## Endpoints

There are 7 endpoints set up within the Banking Api. Basic validation is performed on each request passed to the endpoints. The endpoints are described below:

| URL                                                         | Method | Description                                                                          |
| ----------------------------------------------------------- | ------ | ------------------------------------------------------------------------------------ |
| `/customers`                                                | GET    | Returns all customers at the Bank                                                    |
| `/customers/{customerId}`                                   | GET    | Returns a single customer at the bank that matches the customer identifier           |
| `/customers/{customerId}/accounts`                          | GET    | Returns all accounts for a single customer                                           |
| `/customers/{customerId}/accounts`                          | POST   | Creates a new account for the specified customer                                     |
| `/customers/{customerId}/accounts/{accountId}`              | GET    | Returns a single account that matches the account identifier for a specific customer |
| `/customers/{customerId}/accounts/{accountId}/transactions` | GET    | Returns all the transactions listed under a specific customer account                |
| `/customers/{customerId}/accounts/{accountId}/transactions` | POST   | Creates a new transactions under a specific customer account                         |
