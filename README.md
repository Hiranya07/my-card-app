# my-card-app
## Description
This application contains business logic for accounts/transaction management of the card issue by any financial institution

## How to use(with out docker)
1. Clone this repo using below command: git clone git@github.com:Hiranya07/my-card-app.git
2. cd my-card-app
4. Run go mod tidy
5. execute go run main.go

## Configurations
##### This project use mysql as DB for storing data.The Database name,Database address & Database password are configurable as per the need.
##### Before running the application the use need to visit config.json found under my-card-app directory
##### User need to provide value for DbName,DbAddr & DbPassWord configurations
![image](https://user-images.githubusercontent.com/12723634/164987476-54969a4a-d5d2-452c-84b7-b182bc00e152.png)

## Rest Endpoints
###### Please refer to the swagger file for the rest endpoints:https://github.com/Hiranya07/my-card-app/blob/main/app/swagger.yml





