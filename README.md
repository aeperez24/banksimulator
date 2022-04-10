# banksimulator
Application features:

## Sign up
Endpoint to  create a new account, required fields are:
- Name
- Rut
- email
- password

every fields are required, and it's allowed to create only one account per Rut.

## Sign in
Endpoint to generate a jwt token,  wich provide authorization to consume others endpoints

# Deposit 
endpoint  to add funds  to account (simulating a bank deposit)

## Withdraw
endpoint to reduce  funds from account (simulating a bank withdraw), accounts can`t have negative funds.

## Transfer
endpoint to move funds between accounts, rut of target account and amount  are required  fields, source account can`t end up with negative balance,  it is required to validate that target account is registered in the system.

# Transaction history
endpoint to retrieve all transactions for a consulted account, including withdraws, deposits and transfers (incoming and outcoming) 