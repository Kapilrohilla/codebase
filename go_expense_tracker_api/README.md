# Expense Tracker

## Features

1. Accounts - manage accounts.
2. Expenses - An account can add his expenses.
3. Split - An account can split his expenses to other peoples.

## Start

Create the .env.local file using .env.example

```bash
go mod tidy
```

## Dependencies

1. Fiber - Http framework
2. GoDotEnv - To Load .env's files
3. Gorm - ORM to interact with sql datbases
4. MySql - MySql database driver for go
5. Validator - Pkg for validation
6. GoLang-Jwt - For Tokenization
7. Crypto - To Ecrypt Password, using **bcrypt** algorithm
